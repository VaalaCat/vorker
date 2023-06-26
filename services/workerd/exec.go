package workerd

import (
	"runtime/debug"
	"vorker/common"
	"vorker/entities"
	"vorker/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RunWorkerEndpoint(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		}
	}()

	UID := c.Param("uid")
	if len(UID) == 0 {
		logrus.Errorf("uid is empty, ctx: %v", c)
		common.RespErr(c, common.RespCodeInvalidRequest, common.RespMsgInvalidRequest, nil)
		return
	}

	userID := c.GetUint(common.UIDKey)
	if userID == 0 {
		logrus.Errorf("userID is empty, ctx: %v", c)
		common.RespErr(c, common.RespCodeInvalidRequest, common.RespMsgInvalidRequest, nil)
		return
	}

	worker, err := models.GetWorkerByUID(userID, UID)
	if err != nil || worker == nil {
		logrus.Errorf("get worker by uid failed, err: %v", err)
		common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		return
	}

	resp, err := worker.Run()
	if err != nil {
		logrus.Errorf("run worker failed, err: %v", err)
		common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		return
	}

	logrus.Infof("user: %v, run worker: %+v success, ", userID, UID)
	common.RespOK(c, common.RespMsgOK, entities.RunWorkerResponse{
		Status:  0,
		RunResp: resp,
	})
}
