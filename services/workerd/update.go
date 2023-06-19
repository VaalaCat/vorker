package workerd

import (
	"fmt"
	"voker/common"
	"voker/entities"
	"voker/models"
	"voker/utils/gost"

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
	FillWorkerValue(worker, true, UID)

	workerRecord, err := models.GetWorkerByUID(userID, UID)
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

	gost.DeleteGost(worker.Name)
	gost.AddGost(worker.TunnelID, worker.Name, worker.Port)
	return GenCapnpConfig()
}
