package node

import (
	"vorker/defs"
	"vorker/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func AddEndpoint(c *gin.Context) {
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
