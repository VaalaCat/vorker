package agent

import (
	"vorker/common"
	"vorker/conf"
	"vorker/entities"
	"vorker/models"
	"vorker/rpc"
	"vorker/services/workerd"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SyncEventHandler(c *gin.Context, req *entities.NotifyEventRequest) {
	err := SyncCall()
	if err != nil {
		logrus.WithError(err).Error("sync event handler error")
		common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		return
	}
	if err := workerd.GenCapnpConfig(); err != nil {
		logrus.WithError(err).Error("sync event handler gen capnp config error")
		common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		return
	}

	common.RespOK(c, common.RespMsgOK, nil)
}

func SyncCall() error {
	logrus.Infof("call sync agent")
	workerList, err := rpc.SyncAgent(conf.AppConfigInstance.MasterEndpoint)
	if err != nil {
		return err
	}
	if err := models.SyncWorkers(workerList); err != nil {
		return err
	}
	return workerd.GenCapnpConfig()
}
