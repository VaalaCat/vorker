package workerd

import (
	"voker/common"
	"voker/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func FlushEndpoint(c *gin.Context) {
	UID := c.Param("uid")
	if len(UID) == 0 {
		logrus.Errorf("uid is empty, ctx: %v", c)
		return
	}

	userID := c.GetUint(common.UIDKey)

	if err := Flush(userID, UID); err != nil {
		c.JSON(500, gin.H{"code": 3, "error": err.Error()})
		logrus.Errorf("failed to flush worker, err: %v, ctx: %v", err, c)
		return
	}

	c.JSON(200, gin.H{"code": 0, "message": "success"})
	logrus.Errorf("flush worker success, ctx: %v", c)
}

func FlushAllEndpoint(c *gin.Context) {
	userID := c.GetUint(common.UIDKey)
	workers, err := models.GetAllWorkers(userID)
	if err != nil {
		common.RespErr(c, common.RespCodeInternalError, err.Error(), nil)
		return
	}

	err = nil
	for _, worker := range workers {
		if err = worker.Flush(); err != nil {
			logrus.Errorf("failed to flush worker, err: %v, ctx: %v", err, c)
			continue
		}
	}
	if err != nil {
		common.RespErr(c, common.RespCodeInternalError, err.Error(), nil)
		logrus.Warnf("partial failure, ctx: %v", c)
		return
	}

	common.RespOK(c, "flush worker success", nil)
}

func Flush(userID uint, UID string) error {
	worker, err := models.GetWorkerByUID(userID, UID)
	if err != nil {
		return err
	}
	err = worker.Flush()
	if err != nil {
		return err
	}
	return GenCapnpConfig()
}
