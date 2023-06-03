package workerd

import (
	"fmt"
	"voker/common"
	"voker/entities"
	"voker/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func UpdateEndpoint(c *gin.Context) {
	UID := c.Param("uid")
	if len(UID) == 0 {
		c.JSON(400, gin.H{"code": 1, "error": "uid is empty"})
		logrus.Errorf("uid is empty, ctx: %v", c)
		return
	}

	var worker *entities.Worker
	if err := c.ShouldBindJSON(&worker); err != nil {
		c.JSON(400, gin.H{"code": 1, "error": err.Error()})
		logrus.Errorf("failed to bind json, err: %v, ctx: %v", err, c)
		return
	}

	userID := c.GetUint(common.UIDKey)

	if err := Update(userID, UID, worker); err != nil {
		c.JSON(500, gin.H{"code": 3, "error": err.Error()})
		logrus.Errorf("failed to update worker, err: %v, ctx: %v", err, c)
		return
	}

	c.JSON(200, gin.H{"code": 0, "message": "success"})
	logrus.Errorf("update worker success, ctx: %v", c)
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

	return GenCapnpConfig()
}
