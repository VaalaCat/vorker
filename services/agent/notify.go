package agent

import (
	"vorker/common"
	"vorker/conf"
	"vorker/rpc"

	"github.com/gin-gonic/gin"
)

func NotifyEndpoint(c *gin.Context) {
	err := rpc.SyncAgent(conf.AppConfigInstance.MasterEndpoint)
	if err != nil {
		common.RespErr(c, common.RespCodeInternalError, err.Error(), nil)
		return
	}
	common.RespOK(c, "success", nil)
}
