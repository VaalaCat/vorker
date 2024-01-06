package node

import (
	"runtime/debug"
	"vorker/common"
	"vorker/conf"
	"vorker/defs"
	"vorker/entities"
	"vorker/models"
	"vorker/rpc"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LeaveEndpoint(c *gin.Context) {
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

	if nodename == defs.DefaultNodeName {
		logrus.WithContext(c).Errorf("you cannot leave the default node")
		common.RespOK(c, "you cannot leave the default node", nil)
		return
	}

	if err := models.AdminDeleteNode(node.Node.UID); err != nil {
		logrus.WithContext(c).Errorf("delete node failed, err: %v", err)
		common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		return
	}

	oldWorkers, err := models.AdminGetWorkersByNodeName(nodename)
	if err != nil {
		logrus.WithContext(c).Errorf("get workers failed, err: %v", err)
		common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		return
	}

	nodeMap := make(map[string]*entities.Node)

	for _, w := range oldWorkers {
		assignNode, err := models.GetAssignNode()
		if err == nil {
			w.NodeName = assignNode.GetName()
			w.TunnelID = assignNode.UID
		} else {
			w.NodeName = defs.DefaultNodeName
			w.TunnelID = conf.AppConfigInstance.NodeID
		}
		nodeMap[w.NodeName] = assignNode.Node
		w.Flush()
	}

	for tNodeName, tNode := range nodeMap {
		logrus.Infof("call sync to tNodeName: %s, tNode: %+v", tNodeName, tNode)
		go rpc.EventNotify(tNode, defs.EventSyncWorkers, nil)
	}

	common.RespOK(c, common.RespMsgOK, nil)
}
