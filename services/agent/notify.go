package agent

import (
	"runtime/debug"
	"vorker/common"
	"vorker/conf"
	"vorker/defs"
	"vorker/entities"
	"vorker/rpc"
	"vorker/utils/request"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NotifyEndpoint(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		}
	}()

	req := &entities.NotifyEventRequest{}
	err := request.Bind[*entities.NotifyEventRequest](c, req)
	if err != nil {
		common.RespErr(c, common.RespCodeInvalidRequest, common.RespMsgInvalidRequest, nil)
		return
	}

	switch req.EventName {
	case defs.EventSyncWorkers:
		if err = rpc.SyncAgent(conf.AppConfigInstance.MasterEndpoint); err != nil {
			logrus.Errorf("event: %s error, err: %+v", req.EventName, err)
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
			return
		} else {
			common.RespOK(c, common.RespMsgOK, nil)
			return
		}
	default:
	}

	if err != nil {
		logrus.Errorf("event: %s error, err: %+v", req.EventName, err)
		common.RespErr(c, common.RespCodeInternalError, err.Error(), nil)
		return
	}
	common.RespOK(c, common.RespMsgOK, nil)
}
