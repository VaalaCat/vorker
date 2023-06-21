package node

import (
	"runtime/debug"
	"vorker/common"
	"vorker/defs"
	"vorker/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func AddEndpoint(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		}
	}()
	nodeName := c.GetString(defs.KeyNodeName)

	newNode := &models.Node{
		UID:  uuid.New().String(),
		Name: nodeName,
	}

	if err := newNode.Create(); err != nil {
		logrus.Errorf("failed to create node, err: %v", err)
		c.JSON(defs.CodeInternalError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(defs.CodeSuccess, gin.H{"message": "success", "data": newNode})
}
