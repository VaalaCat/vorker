package node

import (
	"fmt"
	"runtime/debug"
	"vorker/common"
	"vorker/conf"
	"vorker/defs"
	"vorker/models"
	"vorker/rpc"
	"vorker/utils"
	"vorker/utils/request"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetNodeInfoEndpoint(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		}
	}()
	nodeName := c.GetString(defs.KeyNodeName)

	node, err := models.GetNodeByNodeName(nodeName)
	if err != nil {
		logrus.Errorf("failed to get node info, err: %v", err)
		common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		return
	}

	go rpc.EventNotify(node.Node, defs.EventSyncWorkers, nil)
	common.RespOK(c, common.RespMsgOK, node)
}

func UserGetNodesEndpoint(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		}
	}()

	nodes, err := models.AdminGetAllNodes()
	if err != nil {
		logrus.Errorf("failed to get nodes, err: %v", err)
		common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		return
	}
	pingMap := map[string]int{}
	for _, node := range nodes {

		var addr string
		if node.Name == conf.AppConfigInstance.NodeName {
			addr = fmt.Sprintf("http://%s:%d", conf.AppConfigInstance.TunnelHost, conf.AppConfigInstance.APIPort)
		} else {
			addr = fmt.Sprintf("http://%s:%d", conf.AppConfigInstance.TunnelHost, conf.AppConfigInstance.TunnelEntryPort)
		}

		pingMap[node.Name], err = request.Ping(
			addr, utils.NodeHost(node.Name, node.UID))
		if err != nil {
			logrus.Errorf("failed to ping node %s, err: %v", node.Name, err)
			pingMap[node.Name] = 9999
		}
	}

	common.RespOK(c, common.RespMsgOK, gin.H{
		"nodes": nodes,
		"ping":  pingMap,
	})
}
