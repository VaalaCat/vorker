package workerd

import (
	"path/filepath"
	"voker/conf"
	"voker/entities"
	"voker/models"
	"voker/utils"

	"github.com/sirupsen/logrus"
)

// Create creates a new worker in the database and update the workerd capnp config file
func Create(worker *entities.Worker) error {
	if err := (&models.Worker{Worker: worker}).Create(); err != nil {
		logrus.Errorf("failed to create worker, err: %v", err)
		return err
	}
	err := utils.WriteFile(
		filepath.Join(conf.AppConfigInstance.WorkerdDir, worker.UID, worker.Entry),
		string(worker.Code))
	if err != nil {
		return err
	}
	return GenCapnpConfig()
}
