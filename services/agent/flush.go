package agent

import (
	"vorker/common"
	"vorker/defs"
	"vorker/entities"
	"vorker/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func FlushWorkerEventHandler(c *gin.Context, req *entities.NotifyEventRequest) {
	worker, err := entities.ToWorkerEntity(req.Extra[defs.KeyWorkerProto])
	if err != nil {
		logrus.WithError(err).Error("flush worker event handler error")
		common.RespErr(c, common.RespCodeInvalidRequest, common.RespMsgInvalidRequest, nil)
		return
	}

	if err := (&models.Worker{Worker: worker}).Flush(); err != nil {
		logrus.WithError(err).Error("flush worker event handler error")
		common.RespErr(c, common.RespCodeInvalidRequest, common.RespMsgInvalidRequest, nil)
		return
	}

	common.RespOK(c, common.RespMsgOK, nil)
}
