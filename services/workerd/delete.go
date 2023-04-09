package workerd

import (
	"fmt"
	"voker/models"
)

func Delete(UID string) error {
	worker, err := models.GetWorkerByUID(UID)
	if err != nil {
		return err
	}
	if worker == nil {
		return fmt.Errorf("worker not found")
	}
	err = worker.Delete()
	if err != nil {
		return err
	}

	return GenCapnpConfig()
}
