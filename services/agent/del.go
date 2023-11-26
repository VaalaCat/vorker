package agent

import (
	"vorker/common"
	"vorker/conf"
	"vorker/defs"
	"vorker/entities"
	"vorker/exec"
	"vorker/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func DelWorkerEventHandler(c *gin.Context, req *entities.NotifyEventRequest) {
	worker, err := entities.ToWorkerEntity(req.Extra[defs.KeyWorkerProto])
	if err != nil || worker == nil {
		logrus.Errorf("event: %s error, err: %+v", req.EventName, err)
		common.RespErr(c, common.RespCodeInvalidRequest, common.RespMsgInvalidRequest, nil)
		return
	}

	if worker.NodeName == conf.AppConfigInstance.NodeName {
		exec.ExecManager.ExitCmd(worker.GetUID())
	}

	if err := (&models.Worker{Worker: worker}).Delete(); err != nil {
		logrus.Errorf("delete worker error, err: %+v", err)
		common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		return
	}

	logrus.Info("delete worker event handler success")
	common.RespOK(c, common.RespMsgOK, nil)
}
