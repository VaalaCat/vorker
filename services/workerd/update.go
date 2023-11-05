package workerd

import (
	"fmt"
	"vorker/common"
	"vorker/entities"
	"vorker/models"

	"github.com/gin-gonic/gin"
)

func UpdateEndpoint(c *gin.Context) {
	UID := c.Param("uid")
	if len(UID) == 0 {
		common.RespErr(c, common.RespCodeInvalidRequest, "uid is empty", nil)
		return
	}

	var worker *entities.Worker
	if err := c.ShouldBindJSON(&worker); err != nil {
		common.RespErr(c, common.RespCodeInvalidRequest, err.Error(), nil)
		return
	}

	userID := c.GetUint(common.UIDKey)

	if err := Update(userID, UID, worker); err != nil {
		common.RespErr(c, common.RespCodeInternalError, err.Error(), nil)
		return
	}

	common.RespOK(c, "update worker success", nil)
}

func Update(userID uint, UID string, worker *entities.Worker) error {
	FillWorkerValue(worker, true, UID, userID)

	workerRecord, err := models.GetWorkerByUID(userID, UID)
	if err != nil {
		return err
	}
	if worker == nil {
		return fmt.Errorf("worker is nil")
	}

	err = workerRecord.Delete()
	if err != nil {
		return err
	}

	worker.Version = worker.GetVersion() + 1
	newWorker := &models.Worker{Worker: worker}
	err = newWorker.Create()
	if err != nil {
		return err
	}
	return GenCapnpConfig()
}
