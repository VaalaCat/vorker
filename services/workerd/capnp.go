package workerd

import (
	"path/filepath"
	"voker/conf"
	"voker/defs"
	"voker/entities"
	"voker/models"
	"voker/utils"
	"voker/utils/workerd"

	"github.com/sirupsen/logrus"
)

func GenCapnpConfig() error {
	workerRecords, err := models.GetAllWorkers()
	if err != nil {
		logrus.Errorf("failed to get all workers, err: %v", err)
	}

	workerList := &entities.WorkerList{
		Workers: models.Trans2Entities(workerRecords),
	}

	return utils.WriteFile(
		filepath.Join(
			conf.AppConfigInstance.WorkerdDir,
			defs.CapFileName,
		),
		workerd.BuildCapfile(workerList))
}
