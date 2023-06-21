package gost

import (
	"context"
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"runtime/debug"
	"sync"
	"vorker/conf"
	"vorker/entities"
	"vorker/utils/idgen"

	"github.com/judwhite/go-svc"
	"github.com/sirupsen/logrus"
)

var (
	cfgFile       string
	outputFormat  string
	services      stringList
	nodes         stringList
	debugMode     bool
	apiAddr       string
	metricsAddr   string
	wg            sync.WaitGroup
	ret           int
	gostCtxMap    map[int64]context.CancelFunc
	workerGostMap map[string]int64
)

func InitGost() {
	for _, f := range gostCtxMap {
		f()
	}
	gostCtxMap = make(map[int64]context.CancelFunc)
	workerGostMap = make(map[string]int64)
	go func() {
		t := entities.GetTunnel().GetAll()
		p := entities.GetProxy().GetAll()
		r := buildGostPool(t, p)
		for workerName, wargs := range r {
			wg.Add(1)
			ctx, cancel := context.WithCancel(context.Background())
			wid := idgen.GetNextID()
			gostCtxMap[wid] = cancel
			go func(wid int64, wargs stringList) {
				defer wg.Done()
				defer cancel()
				worker(wid, wargs, &ctx, &ret)
			}(wid, wargs)
			workerGostMap[workerName] = wid
		}
		wg.Wait()
		for _, f := range gostCtxMap {
			f()
		}
	}()
}

func init() {
	InitGost()
}

func AddGost(tunnelID string, workerName string, workerPort int32) int64 {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
		}
	}()
	r := buildGostArgs(conf.AppConfigInstance.TunnelScheme, "127.0.0.1", workerPort,
		conf.AppConfigInstance.TunnelRelayEndpoint, tunnelID)
	wid := idgen.GetNextID()
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	gostCtxMap[wid] = cancel
	go func(wid int64, wargs stringList) {
		defer wg.Done()
		defer cancel()
		worker(wid, wargs, &ctx, &ret)
	}(wid, r)
	workerGostMap[workerName] = wid
	return wid
}

func DeleteGost(workerName string) {
	wid := workerGostMap[workerName]
	gostCtxMap[wid]()
	delete(gostCtxMap, wid)
	delete(workerGostMap, workerName)
}

func worker(id int64, args []string, ctx *context.Context, ret *int) {
	cmd := exec.CommandContext(*ctx, conf.AppConfigInstance.GostBinPath, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), fmt.Sprintf("_GOST_ID=%d", id))

	err := cmd.Run()
	if err != nil {
		logrus.Errorf("gost worker error: %v", err)
	}
	if cmd.ProcessState.Exited() {
		*ret = cmd.ProcessState.ExitCode()
	}
}

func Run() {
	if conf.AppConfigInstance.RunMode == "master" {
		p := &program{}
		if err := svc.Run(p); err != nil {
			log.Fatal(err)
		}
	}
}