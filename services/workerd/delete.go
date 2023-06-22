package workerd

import (
	"fmt"
	"runtime/debug"
	"vorker/common"
	"vorker/entities"
	"vorker/models"
	"vorker/utils/gost"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func DeleteEndpoint(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		}
	}()
	UID := c.Param("uid")
	if len(UID) == 0 {
		common.RespErr(c, common.RespCodeInvalidRequest, "uid is empty", nil)
		return
	}

	userID := c.GetUint(common.UIDKey)
	if err := Delete(userID, UID); err != nil {
		common.RespErr(c, common.RespCodeInternalError, err.Error(), nil)
		return
	}

	go func() {
		worker, err := models.GetWorkerByUID(userID, UID)
		if err != nil {
			common.RespErr(c, common.RespCodeInternalError, err.Error(), nil)
			return
		}
		SyncAgent(worker.Worker)
	}()
	common.RespOK(c, "delete worker success", nil)
}

func Delete(userID uint, UID string) error {
	worker, err := models.GetWorkerByUID(userID, UID)
	if err != nil {
		return err
	}
	if worker == nil {
		return fmt.Errorf("worker not found")
	}

	if err = worker.Delete(); err != nil {
		return err
	}

	if err = GenCapnpConfig(); err != nil {
		return err
	}

	worker.DeleteFile()
	entities.GetTunnel().DeleteTunnel(worker.Worker)
	gost.DeleteGost(worker.Name)
	return nil
}
