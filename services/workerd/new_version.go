package workerd

import (
	"vorker/common"
	"vorker/dao"
	"vorker/models"
	"vorker/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewVersionEndpoint(c *gin.Context) {
	workerID := c.Param("workerId")
	fileID := c.Param("fileId")
	userID := c.GetUint(common.UIDKey)

	if len(workerID) == 0 || len(fileID) == 0 {
		common.RespErr(c, common.RespCodeInvalidRequest, "workerId or fileId is empty", nil)
		return
	}

	worker, err := models.GetWorkerByUID(userID, workerID)
	if err != nil || worker == nil {
		common.RespErr(c, common.RespCodeInvalidRequest, "worker not found", nil)
		return
	}

	file, err := dao.GetFileByUID(c, userID, fileID)
	if err != nil || file == nil {
		common.RespErr(c, common.RespCodeInvalidRequest, "file not found", nil)
		return
	}

	versionID := utils.GenerateUID()

	if err := dao.NewWorkerVersion(&models.WorkerVersion{
		UID:      versionID,
		WorkerID: workerID,
		FileID:   fileID,
		Name:     utils.NewCodeName(0),
	}); err != nil {
		logrus.WithError(err).Error("new version error")
		common.RespErr(c, common.RespCodeInternalError, "new version error", nil)
		return
	}

	worker.ActiveVersionID = versionID

	if err := UpdateWorker(userID, workerID, worker.Worker); err != nil {
		common.RespErr(c, common.RespCodeInternalError, "update worker error", nil)
		return
	}

	common.RespOK(c, "new version success", nil)
}
