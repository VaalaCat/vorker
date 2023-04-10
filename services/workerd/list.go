package workerd

import (
	"strconv"
	"voker/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetWorkersEndpoint(c *gin.Context) {
	offsetStr := c.Param("offset")
	if len(offsetStr) == 0 {
		c.JSON(400, gin.H{"code": 1, "error": "offset is empty"})
		return
	}
	limitStr := c.Param("limit")
	if len(limitStr) == 0 {
		c.JSON(400, gin.H{"code": 1, "error": "limit is empty"})
		return
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		c.JSON(400, gin.H{"code": 1, "error": "offset is invalid"})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(400, gin.H{"code": 1, "error": "limit is invalid"})
		return
	}

	workers, err := models.GetWorkers(offset, limit)
	if err != nil {
		c.JSON(500, gin.H{"code": 3, "error": err.Error()})
		logrus.Errorf("failed to get worker, err: %v, ctx: %v", err, c)
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "success", "data": models.Trans2Entities(workers)})
	logrus.Infof("get worker success, ctx: %v", c)
}

func GetAllWorkersEndpoint(c *gin.Context) {
	workers, err := models.GetAllWorkers()
	if err != nil {
		c.JSON(500, gin.H{"code": 3, "error": err.Error()})
		logrus.Errorf("failed to get all workers, err: %v, ctx: %v", err, c)
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "success", "data": models.Trans2Entities(workers)})
	logrus.Infof("get all workers success, ctx: %v", c)
}

func GetWorkerEndpoint(c *gin.Context) {
	uid := c.Param("uid")
	if len(uid) == 0 {
		c.JSON(400, gin.H{"code": 1, "error": "uid is empty"})
		return
	}
	worker, err := models.GetWorkerByUID(uid)
	if err != nil {
		c.JSON(500, gin.H{"code": 3, "error": err.Error()})
		logrus.Errorf("failed to get worker, err: %v, ctx: %v", err, c)
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "success", "data": models.Trans2Entities([]*models.Worker{worker})})
	logrus.Infof("get worker success, ctx: %v", c)
}
