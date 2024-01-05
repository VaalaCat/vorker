package models

import (
	"vorker/conf"
	"vorker/exec"
	"vorker/utils"

	"github.com/sirupsen/logrus"
)

func NodeWorkersInit() {
	workerRecords, err := AdminGetWorkersByNodeName(conf.AppConfigInstance.NodeName)
	if err != nil {
		logrus.Errorf("init failed to get all workers, err: %v", err)
	}
	logrus.Infof("this node will init %d workers", len(workerRecords))
	for _, worker := range workerRecords {
		worker.Flush()
		utils.GenWorkerConfig(worker.Worker)
		exec.ExecManager.RunCmd(worker.GetUID(), []string{})
	}
}
