package workerd

import (
	"runtime/debug"
	"strconv"
	"vorker/common"
	"vorker/defs"
	"vorker/entities"
	"vorker/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetWorkersEndpoint(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		}
	}()
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
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		}
	}()
	userID := c.GetUint(common.UIDKey)
	workers, err := models.GetAllWorkers(userID)
	if err != nil {
		common.RespErr(c, common.RespCodeInternalError, err.Error(), nil)
		return
	}

	common.RespOK(c, "get all workers success", models.Trans2Entities(workers))
}

func GetWorkerEndpoint(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		}
	}()
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

func AgentSyncWorkers(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		}
	}()
	req := &defs.AgentSyncWorkersReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		common.RespErr(c, defs.CodeInvalidRequest, err.Error(), nil)
		return
	}

	nodeName := c.GetString(defs.KeyNodeName)
	// get node's workerlist
	workers, err := models.AdminGetWorkersByNodeName(nodeName)
	if err != nil {
		common.RespErr(c, defs.CodeInternalError, err.Error(), nil)
		return
	}

	// build response
	// TODO: chunk loading
	resp := &defs.AgentSyncWorkersResp{
		WorkerList: &entities.WorkerList{
			NodeName: nodeName,
			Workers:  models.Trans2Entities(workers),
		},
	}
	common.RespOK(c, "sync workers success", resp)
}
