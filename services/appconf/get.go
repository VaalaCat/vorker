package appconf

import (
	"net/http"
	"vorker/common"
	"vorker/conf"
	"vorker/models"

	"github.com/gin-gonic/gin"
)

func GetEndpoint(c *gin.Context) {
	num, err := models.AdminGetUserNumber()
	if err != nil {
		common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"WorkerURLSuffix": conf.AppConfigInstance.WorkerURLSuffix,
			"Scheme":          conf.AppConfigInstance.Scheme,
			"EnableRegister":  conf.AppConfigInstance.EnableRegister || num == 0,
		},
	})
}
