package workerd

import (
	"strconv"
	"voker/common"
	"voker/models"

	"github.com/gin-gonic/gin"
)

func GetWorkersEndpoint(c *gin.Context) {
	offsetStr := c.Param("offset")
	if len(offsetStr) == 0 {
		common.RespErr(c, common.RespCodeInvalidRequest, "offset is empty", nil)
		return
	}
	limitStr := c.Param("limit")
	if len(limitStr) == 0 {
		common.RespErr(c, common.RespCodeInvalidRequest, "limit is empty", nil)
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		common.RespErr(c, common.RespCodeInvalidRequest, "offset is invalid", nil)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		common.RespErr(c, common.RespCodeInvalidRequest, "limit is invalid", nil)
		return
	}
	userID := c.GetUint(common.UIDKey)

	workers, err := models.GetWorkers(userID, offset, limit)
	if err != nil {
		common.RespErr(c, common.RespCodeInternalError, err.Error(), nil)
		return
	}

	common.RespOK(c, "get worker success", models.Trans2Entities(workers))
}

func GetAllWorkersEndpoint(c *gin.Context) {
	userID := c.GetUint(common.UIDKey)
	workers, err := models.GetAllWorkers(userID)
	if err != nil {
		common.RespErr(c, common.RespCodeInternalError, err.Error(), nil)
		return
	}

	common.RespOK(c, "get all workers success", models.Trans2Entities(workers))
}

func GetWorkerEndpoint(c *gin.Context) {
	userID := c.GetUint(common.UIDKey)
	uid := c.Param("uid")
	if len(uid) == 0 {
		common.RespErr(c, common.RespCodeInvalidRequest, "uid is empty", nil)
		return
	}
	worker, err := models.GetWorkerByUID(userID, uid)
	if err != nil {
		common.RespErr(c, common.RespCodeInternalError, err.Error(), nil)
		return
	}
	common.RespOK(c, "get workers success", models.Trans2Entities([]*models.Worker{worker}))
}
