package agent

import (
	"vorker/common"
	"vorker/entities"
	"vorker/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type EventRouter interface {
	RegisteHandler(eventName string, handler func(c *gin.Context, req *entities.NotifyEventRequest))
	Handle(c *gin.Context, req *entities.NotifyEventRequest)
}

type EventRouterImpl struct {
	handlers map[string]func(c *gin.Context, req *entities.NotifyEventRequest)
}

func (e *EventRouterImpl) RegisteHandler(eventName string, handler func(c *gin.Context, req *entities.NotifyEventRequest)) {
	if e.handlers == nil {
		e.handlers = make(map[string]func(c *gin.Context, req *entities.NotifyEventRequest))
	}
	e.handlers[eventName] = handler
}

func (e *EventRouterImpl) Handle(c *gin.Context, req *entities.NotifyEventRequest) {
	if handler, ok := e.handlers[req.EventName]; ok {
		logrus.Infof("handle event: %s, extra key: %+v", req.EventName, utils.GetKey(req.Extra))
		handler(c, req)
	} else {
		logrus.Errorf("handle event: %s, extra key: %+v", req.EventName, utils.GetKey(req.Extra))
		common.RespErr(c, common.RespCodeInvalidRequest, common.RespMsgInvalidRequest, nil)
	}
}

var (
	EventRouterImplInstance = &EventRouterImpl{}
)
