package agent

import (
	"runtime/debug"
	"vorker/common"
	"vorker/conf"
	"vorker/rpc"

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
	err := rpc.SyncAgent(conf.AppConfigInstance.MasterEndpoint)
	if err != nil {
		common.RespErr(c, common.RespCodeInternalError, err.Error(), nil)
		return
	}
	common.RespOK(c, "success", nil)
}
