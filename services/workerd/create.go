package workerd

import (
	"voker/common"
	"voker/entities"
	"voker/models"
	"voker/utils/gost"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateEndpoint(c *gin.Context) {
	worker := &entities.Worker{}

	if err := c.BindJSON(worker); err != nil {
		common.RespErr(c, common.RespCodeInvalidRequest, err.Error(), nil)
		return
	}

	if !isCreateParamValidate() {
		common.RespErr(c, common.RespCodeInvalidRequest, common.RespMsgInvalidRequest, nil)
		return
	}
	userID := c.GetUint(common.UIDKey)

	if err := Create(userID, worker); err != nil {
		common.RespErr(c, common.RespCodeInternalError, err.Error(), nil)
		return
	}

	common.RespOK(c, "create worker success", nil)
}

// Create creates a new worker in the database and update the workerd capnp config file
func Create(userID uint, worker *entities.Worker) error {
	FillWorkerValue(worker, false, "")

	if err := (&models.Worker{Worker: worker, UserID: userID}).Create(); err != nil {
		logrus.Errorf("failed to create worker, err: %v", err)
		return err
	}

	gost.AddGost(worker.TunnelID, worker.Name, worker.Port)
	return GenCapnpConfig()
}

func isCreateParamValidate() bool {
	// TODO: validate the create params
	return true
}
