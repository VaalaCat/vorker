package workerd

import (
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

	if err := Flush(UID); err != nil {
		c.JSON(500, gin.H{"code": 3, "error": err.Error()})
		logrus.Errorf("failed to flush worker, err: %v, ctx: %v", err, c)
		return
	}

	c.JSON(200, gin.H{"code": 0, "message": "success"})
	logrus.Errorf("flush worker success, ctx: %v", c)
}

func FlushAllEndpoint(c *gin.Context) {
	workers, err := models.GetAllWorkers()
	if err != nil {
		c.JSON(500, gin.H{"code": 3, "error": err.Error()})
		logrus.Errorf("failed to get all workers, err: %v, ctx: %v", err, c)
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
		c.JSON(200, gin.H{"code": 10, "message": "partial failure"})
		logrus.Warnf("partial failure, ctx: %v", c)
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "success"})
	logrus.Infof("flush all workers success, ctx: %v", c)
}

func Flush(UID string) error {
	worker, err := models.GetWorkerByUID(UID)
	if err != nil {
		return err
	}
	err = worker.Flush()
	if err != nil {
		return err
	}
	return GenCapnpConfig()
}
