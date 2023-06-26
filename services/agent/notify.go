package agent

import (
	"runtime/debug"
	"vorker/common"
	"vorker/defs"
	"vorker/entities"
	"vorker/utils/request"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	EventRouterImplInstance.RegisteHandler(defs.EventSyncWorkers, SyncEventHandler)
	EventRouterImplInstance.RegisteHandler(defs.EventAddWorker, AddWorkerEventHandler)
	EventRouterImplInstance.RegisteHandler(defs.EventDeleteWorker, DelWorkerEventHandler)
}

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
		logrus.Errorf("event: %s error, err: %+v", req.EventName, err)
		common.RespErr(c, common.RespCodeInvalidRequest, common.RespMsgInvalidRequest, nil)
		return
	}

	EventRouterImplInstance.Handle(c, req)
}
