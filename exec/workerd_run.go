package exec

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
	"vorker/conf"
	"vorker/defs"

	"github.com/sirupsen/logrus"
)

type execManager struct {
	//用于外层循坏的退出
	signMap *defs.SyncMap[string, bool]
	//用于执行cancel函数
	chanMap *defs.SyncMap[string, chan struct{}]
}

var ExecManager *execManager

func init() {
	ExecManager = &execManager{
		signMap: new(defs.SyncMap[string, bool]),
		chanMap: new(defs.SyncMap[string, chan struct{}]),
	}
}

func (m *execManager) RunCmd(uid string, argv []string) {
	if _, ok := m.chanMap.Get(uid); ok {
		logrus.Warnf("workerd %s is already running!", uid)
		return
	}

	c := make(chan struct{})
	m.chanMap.Set(uid, c)

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context, uid string, argv []string, m *execManager) {
		defer func(uid string, m *execManager) {
			m.signMap.Delete(uid)
		}(uid, m)

		logrus.Infof("workerd %s running!", uid)
		workerdDir := filepath.Join(
			conf.AppConfigInstance.WorkerdDir,
			defs.WorkerInfoPath,
			uid,
		)

		for {
			args := []string{"serve",
				filepath.Join(workerdDir, defs.CapFileName),
				"--watch", "--verbose"}
			args = append(args, argv...)
			cmd := exec.CommandContext(ctx, conf.AppConfigInstance.WorkerdBinPath, args...)
			cmd.Dir = workerdDir
			cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: false}
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()

			if exit, ok := m.signMap.Get(uid); ok && exit {
				return
			}
			time.Sleep(3 * time.Second)
		}
	}(ctx, uid, argv, m)

	go func(cancel context.CancelFunc, uid string, m *execManager) {
		defer func(uid string, m *execManager) {
			m.chanMap.Delete(uid)
		}(uid, m)

		if channel, ok := m.chanMap.Get(uid); ok {
			<-channel
			m.signMap.Set(uid, true)
			cancel()
			return
		} else {
			logrus.Errorf("workerd %s is not running!", uid)
			return
		}
	}(cancel, uid, m)
}

func (m *execManager) ExitCmd(uid string) {
	if channel, ok := m.chanMap.Get(uid); ok {
		channel <- struct{}{}
	}
}

func (m *execManager) ExitAllCmd() {
	for uid := range m.chanMap.ToMap() {
		m.ExitCmd(uid)
	}
}
