package workerd

import (
	"fmt"
	"voker/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func DeleteEndpoint(c *gin.Context) {
	UID := c.Param("uid")
	if len(UID) == 0 {
		c.JSON(400, gin.H{"code": 1, "error": "uid is empty"})
		return
	}

	if err := Delete(UID); err != nil {
		c.JSON(500, gin.H{"code": 3, "error": err.Error()})
		logrus.Errorf("failed to delete worker, err: %v, ctx: %v", err, c)
		return
	}

	c.JSON(200, gin.H{"code": 0, "message": "success"})
}

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
