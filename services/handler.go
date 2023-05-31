package services

import (
	"fmt"
	"voker/conf"
	"voker/services/appconf"
	proxyService "voker/services/proxy"
	"voker/services/workerd"
	"voker/utils"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	proxy  *gin.Engine
)

func init() {
	router = gin.Default()
	proxy = gin.Default()
	router.Use(utils.GinMiddleware("http://localhost:3000"))
	api := router.Group("/api")
	{
		api.GET("/worker/:uid", workerd.GetWorkerEndpoint)
		api.POST("/worker/create", workerd.CreateEndpoint)
		api.PATCH("/worker/:uid", workerd.UpdateEndpoint)
		api.DELETE("/worker/:uid", workerd.DeleteEndpoint)
		api.GET("/worker/flush/:uid", workerd.FlushEndpoint)
		api.GET("/workers/flush", workerd.FlushAllEndpoint)
		api.GET("/workers", workerd.GetAllWorkersEndpoint)
		api.GET("/workers/:offset/:limit", workerd.GetWorkersEndpoint)
		api.GET("/vorker/config", appconf.GetEndpoint)
	}

	proxy.Any("/*proxyPath", proxyService.Endpoint)
}

func Run() {
	WorkerdRun(conf.AppConfigInstance.WorkerdDir, []string{})
	go proxy.Run(fmt.Sprintf("%v:%d", conf.AppConfigInstance.ListenAddr, conf.AppConfigInstance.WorkerPort))
	router.Run(fmt.Sprintf("%v:%d", conf.AppConfigInstance.ListenAddr, conf.AppConfigInstance.APIPort))
}
