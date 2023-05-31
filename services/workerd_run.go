package services

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
	"voker/conf"
	"voker/defs"

	"github.com/sirupsen/logrus"
)

func WorkerdRun(workerdDir string, argv []string) {
	go func() {
		logrus.Info("workerd running!")
		for {
			args := []string{"serve",
				filepath.Join(workerdDir, defs.CapFileName),
				"--watch", "--verbose"}
			args = append(args, argv...)
			cmd := exec.Command(conf.AppConfigInstance.WorkerdBinPath, args...)
			cmd.Dir = workerdDir
			cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: false}
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
			time.Sleep(3 * time.Second)
		}
	}()
}
