package services

import (
	"fmt"
	"voker/conf"
	proxyService "voker/services/proxy"
	"voker/services/workerd"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	proxy  *gin.Engine
)

func init() {
	router = gin.Default()
	proxy = gin.Default()
	api := router.Group("/api")
	{
		api.POST("/workers/create", workerd.CreateEndpoint)
		api.PATCH("/workers/:uid", workerd.UpdateEndpoint)
		api.DELETE("/workers/:uid", workerd.DeleteEndpoint)
		api.GET("/workers/flush/:uid", workerd.FlushEndpoint)
		api.GET("/workers/flush", workerd.FlushAllEndpoint)
		api.GET("/workers", workerd.GetAllWorkersEndpoint)
		api.GET("/workers/:offset/:limit", workerd.GetWorkersEndpoint)
	}

	proxy.Any("/*proxyPath", proxyService.Endpoint)
}

func Run() {
	WorkerdRun(conf.AppConfigInstance.WorkerdDir, []string{})
	go proxy.Run(fmt.Sprintf("%v:%d", conf.AppConfigInstance.ListenAddr, conf.AppConfigInstance.ProxyPort))
	router.Run(fmt.Sprintf("%v:%d", conf.AppConfigInstance.ListenAddr, conf.AppConfigInstance.APIPort))
}
