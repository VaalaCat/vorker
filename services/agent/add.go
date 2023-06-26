package agent

import (
	"vorker/common"
	"vorker/defs"
	"vorker/entities"
	"vorker/models"
	"vorker/services/workerd"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AddWorkerEventHandler(c *gin.Context, req *entities.NotifyEventRequest) {
	worker, err := entities.ToWorkerEntity(req.Extra[defs.KeyWorkerProto])
	if err != nil {
		logrus.WithError(err).Error("add worker event handler error")
		common.RespErr(c, common.RespCodeInvalidRequest, common.RespMsgInvalidRequest, nil)
		return
	}

	if err := (&models.Worker{Worker: worker}).Create(); err != nil {
		logrus.WithError(err).Error("add worker event handler error")
		common.RespErr(c, common.RespCodeInvalidRequest, common.RespMsgInvalidRequest, nil)
		return
	}

	if err := workerd.GenCapnpConfig(); err != nil {
		logrus.WithError(err).Error("add worker event handler error")
		common.RespErr(c, common.RespCodeInvalidRequest, common.RespMsgInvalidRequest, nil)
		return
	}

	logrus.Info("add worker event handler success")
	common.RespOK(c, common.RespMsgOK, nil)
}
