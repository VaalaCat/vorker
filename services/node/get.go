package node

import (
	"vorker/common"
	"vorker/defs"
	"vorker/models"
	"vorker/rpc"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetNodeInfoEndpoint(c *gin.Context) {
	nodeName := c.GetString(defs.KeyNodeName)

	node, err := models.GetNodeByNodeName(nodeName)
	if err != nil {
		logrus.Errorf("failed to get node info, err: %v", err)
		common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		return
	}

	go rpc.EventNotify(node, defs.EventSyncWorkers)
	common.RespOK(c, common.RespMsgOK, node)
}
