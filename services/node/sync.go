package node

import (
	"runtime/debug"
	"vorker/common"
	"vorker/defs"
	"vorker/models"
	"vorker/rpc"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SyncNodeEndpoint(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		}
	}()

	nodename := c.Param("nodename")
	if len(nodename) == 0 {
		logrus.WithContext(c).Errorf("nodename is empty")
		common.RespErr(c, common.RespCodeInvalidRequest, common.RespMsgInvalidRequest, nil)
		return
	}

	node, err := models.GetNodeByNodeName(nodename)
	if err != nil || node == nil {
		logrus.WithContext(c).Errorf("get node by nodename failed, err: %v", err)
		common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		return
	}

	if err := rpc.EventNotify(node.Node, defs.EventSyncWorkers, nil); err != nil {
		logrus.WithContext(c).Errorf("event notify failed, err: %v", err)
		common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		return
	}

	logrus.WithContext(c).Infof("sync node: %+v success, ", nodename)
	common.RespOK(c, common.RespMsgOK, nil)
}
