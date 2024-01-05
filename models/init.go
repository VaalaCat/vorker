package models

import (
	"vorker/conf"
	"vorker/exec"

	"github.com/sirupsen/logrus"
)

func NodeWorkersInit() {
	workerRecords, err := AdminGetWorkersByNodeName(conf.AppConfigInstance.NodeName)
	if err != nil {
		logrus.Errorf("init failed to get all workers, err: %v", err)
	}
	for _, worker := range workerRecords {
		exec.ExecManager.RunCmd(worker.GetUID(), []string{})
	}
}
