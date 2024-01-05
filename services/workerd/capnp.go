package workerd

import (
	"errors"
	"path/filepath"
	"vorker/conf"
	"vorker/defs"
	"vorker/entities"
	"vorker/models"
	"vorker/utils"

	"github.com/sirupsen/logrus"
)

func GenCapnpConfig() error {
	workerRecords, err := models.AdminGetWorkersByNodeName(conf.AppConfigInstance.NodeName)
	if err != nil {
		logrus.Errorf("failed to get all workers, err: %v", err)
	}

	workerList := models.Trans2Entities(workerRecords)

	var hasError bool
	for _, worker := range workerList {
		w := &models.Worker{Worker: worker}
		if err := w.Flush(); err != nil {
			logrus.WithError(err).Errorf("failed to flush worker, worker is: %+v", worker)
			hasError = true
			continue
		}
		fileMap := utils.BuildCapfile([]*entities.Worker{w.Worker})

		if fileContent, ok := fileMap[worker.GetUID()]; ok {
			err := utils.WriteFile(
				filepath.Join(
					conf.AppConfigInstance.WorkerdDir,
					defs.WorkerInfoPath,
					worker.GetUID(),
					defs.CapFileName,
				), fileContent)
			if err != nil {
				hasError = true
			}
		}
	}

	if hasError {
		return errors.New("GenCapnpConfig has error")
	}
	return nil
}
