package agent

import (
	"voker/common"
	"voker/conf"
	"voker/rpc"

	"github.com/gin-gonic/gin"
)

func SyncNotifyEndpoint(c *gin.Context) {
	err := rpc.SyncAgent(conf.AppConfigInstance.MasterEndpoint)
	if err != nil {
		common.RespErr(c, common.RespCodeInternalError, err.Error(), nil)
		return
	}
	common.RespOK(c, "success", nil)
}
