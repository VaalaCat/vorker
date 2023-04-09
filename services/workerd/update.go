package workerd

import (
	"fmt"
	"path/filepath"
	"voker/conf"
	"voker/entities"
	"voker/models"
	"voker/utils"
)

func Update(UID string, worker *entities.Worker) error {
	workerRecord, err := models.GetWorkerByUID(UID)
	if err != nil {
		return err
	}
	if worker == nil {
		return fmt.Errorf("worker is nil")
	}

	workerRecord.Worker = worker
	err = workerRecord.Update()
	if err != nil {
		return err
	}

	err = utils.WriteFile(
		filepath.Join(conf.AppConfigInstance.WorkerdDir, worker.UID, worker.Entry),
		string(worker.Code))
	if err != nil {
		return err
	}

	return GenCapnpConfig()
}
