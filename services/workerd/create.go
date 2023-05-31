package workerd

import (
	"voker/entities"
	"voker/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateEndpoint(c *gin.Context) {
	worker := &entities.Worker{}

	if err := c.BindJSON(worker); err != nil {
		c.JSON(400, gin.H{"code": 1, "error": err.Error()})
		logrus.Errorf("failed to bind json, err: %v, ctx: %v", err, c)
		return
	}

	if !isCreateParamValidate() {
		c.JSON(400, gin.H{"code": 1, "error": "create endpoint params is not validate"})
		logrus.Errorf("create endpoint params is not validate, ctx: %v", c)
		return
	}

	if err := Create(worker); err != nil {
		c.JSON(500, gin.H{"code": 3, "error": err.Error()})
		logrus.Errorf("failed to create worker, err: %v, ctx: %v", err, c)
		return
	}

	c.JSON(200, gin.H{"code": 0, "message": "create worker success", "uid": worker.UID})
	logrus.Infof("create worker success, ctx: %v", c)
}

// Create creates a new worker in the database and update the workerd capnp config file
func Create(worker *entities.Worker) error {
	FillWorkerValue(worker, false, "")

	if err := (&models.Worker{Worker: worker}).Create(); err != nil {
		logrus.Errorf("failed to create worker, err: %v", err)
		return err
	}

	return GenCapnpConfig()
}

func isCreateParamValidate() bool {
	// TODO: validate the create params
	return true
}
