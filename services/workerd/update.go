package workerd

import (
	"fmt"
	"vorker/common"
	"vorker/conf"
	"vorker/entities"
	"vorker/exec"
	"vorker/models"
	"vorker/utils"

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

	if err := UpdateWorker(userID, UID, worker); err != nil {
		common.RespErr(c, common.RespCodeInternalError, err.Error(), nil)
		return
	}

	common.RespOK(c, "update worker success", nil)
}

func UpdateWorker(userID uint, UID string, worker *entities.Worker) error {
	FillWorkerValue(worker, true, UID, userID)

	workerRecord, err := models.GetWorkerByUID(userID, UID)
	if err != nil {
		return err
	}
	if worker == nil {
		return fmt.Errorf("worker is nil")
	}

	curNodeName := conf.AppConfigInstance.NodeName

	if workerRecord.NodeName == curNodeName && worker.NodeName != curNodeName {
		exec.ExecManager.ExitCmd(workerRecord.GetUID())
	}

	err = workerRecord.Delete()
	if err != nil {
		return err
	}

	newWorker := &models.Worker{Worker: worker}
	err = newWorker.Create()
	if err != nil {
		return err
	}

	if worker.NodeName == curNodeName {
		err := utils.GenWorkerConfig(newWorker.ToEntity())
		if err != nil {
			return err
		}
		exec.ExecManager.RunCmd(worker.GetUID(), []string{})
	}
	return nil
}
