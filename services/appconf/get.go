package appconf

import (
	"net/http"
	"voker/conf"

	"github.com/gin-gonic/gin"
)

func GetEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"WorkerURLSuffix": conf.AppConfigInstance.WorkerURLSuffix,
			"Scheme":             conf.AppConfigInstance.Scheme,
		},
	})
}
