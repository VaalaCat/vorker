package workerd

import (
	"fmt"
	"vorker/common"
	"vorker/models"
	"vorker/utils/gost"

	"github.com/gin-gonic/gin"
)

func DeleteEndpoint(c *gin.Context) {
	UID := c.Param("uid")
	if len(UID) == 0 {
		c.JSON(400, gin.H{"code": 1, "error": "uid is empty"})
		return
	}

	userID := c.GetUint(common.UIDKey)
	if err := Delete(userID, UID); err != nil {
		common.RespErr(c, common.RespCodeInternalError, err.Error(), nil)
		return
	}

	common.RespOK(c, "delete worker success", nil)
}

func Delete(userID uint, UID string) error {
	worker, err := models.GetWorkerByUID(userID, UID)
	if err != nil {
		return err
	}
	if worker == nil {
		return fmt.Errorf("worker not found")
	}
	err = worker.Delete()
	if err != nil {
		return err
	}
	gost.DeleteGost(worker.Name)

	return GenCapnpConfig()
}
