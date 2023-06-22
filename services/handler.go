package services

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"vorker/authz"
	"vorker/conf"
	"vorker/rpc"
	"vorker/services/agent"
	"vorker/services/appconf"
	"vorker/services/auth"
	"vorker/services/node"
	proxyService "vorker/services/proxy"
	"vorker/services/tunnel"
	"vorker/services/workerd"
	"vorker/utils"
	"vorker/utils/gost"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
	proxy  *gin.Engine
)

func init() {
	router = gin.Default()
	proxy = gin.Default()
	router.Use(utils.CORSMiddlewaire(
		fmt.Sprintf("%v://%v", conf.AppConfigInstance.Scheme, conf.AppConfigInstance.CookieDomain),
	))
	api := router.Group("/api")
	{
		workerApi := api.Group("/worker", authz.JWTMiddleware())
		{
			workerApi.GET("/:uid", workerd.GetWorkerEndpoint)
			workerApi.GET("/flush/:uid", workerd.FlushEndpoint)
			workerApi.POST("/create", workerd.CreateEndpoint)
			workerApi.PATCH("/:uid", workerd.UpdateEndpoint)
			workerApi.DELETE("/:uid", workerd.DeleteEndpoint)
		}
		workersApi := api.Group("/workers", authz.JWTMiddleware())
		{
			workersApi.GET("/flush", workerd.FlushAllEndpoint)
			workersApi.GET("/:offset/:limit", workerd.GetWorkersEndpoint)
		}
		userApi := api.Group("/user", authz.JWTMiddleware())
		{
			userApi.GET("/info", auth.GetUserEndpoint)
		}
		agentAPI := api.Group("/agent")
		{
			if conf.AppConfigInstance.RunMode == "master" {
				agentAPI.POST("/sync", authz.AgentAuthz(), workerd.AgentSyncWorkers)
				agentAPI.GET("/ingress", tunnel.GetIngressConf)
				agentAPI.POST("/add", authz.AgentAuthz(), node.AddEndpoint)
				agentAPI.GET("/nodeinfo", authz.AgentAuthz(), node.GetNodeInfoEndpoint)
			} else {
				agentAPI.GET("/notify", authz.AgentAuthz(), agent.NotifyEndpoint)
				gost.AddGost(conf.AppConfigInstance.NodeID,
					fmt.Sprintf("%s%s", conf.AppConfigInstance.NodeName, conf.AppConfigInstance.NodeID),
					int32(conf.AppConfigInstance.APIPort))
			}
		}
		api.GET("/allworkers", authz.JWTMiddleware(), workerd.GetAllWorkersEndpoint)
		api.GET("/vorker/config", appconf.GetEndpoint)
		api.POST("/auth/register", auth.RegisterEndpoint)
		api.POST("/auth/login", auth.LoginEndpoint)
		api.GET("/auth/logout", authz.JWTMiddleware(), auth.LogoutEndpoint)
	}

	proxy.Any("/*proxyPath", proxyService.Endpoint)
}

func Run(f embed.FS) {
	WorkerdRun(conf.AppConfigInstance.WorkerdDir, []string{})
	go proxy.Run(fmt.Sprintf("%v:%d", conf.AppConfigInstance.ListenAddr, conf.AppConfigInstance.WorkerPort))
	{
		fp, err := fs.Sub(f, "www/out")
		if err != nil {
			logrus.Panic(err)
		}
		router.StaticFileFS("/404", "404.html", http.FS(fp))
		router.StaticFileFS("/login", "login.html", http.FS(fp))
		router.StaticFileFS("/admin", "admin.html", http.FS(fp))
		router.StaticFileFS("/register", "register.html", http.FS(fp))
		router.StaticFileFS("/worker", "worker.html", http.FS(fp))
		router.StaticFileFS("/index", "index.html", http.FS(fp))
		router.NoRoute(func(c *gin.Context) {
			c.FileFromFS(c.Request.URL.Path, http.FS(fp))
		})
	}
	go gost.Run()
	if conf.AppConfigInstance.RunMode == "agent" {
		self, err := rpc.GetNode(conf.AppConfigInstance.MasterEndpoint)
		if err != nil || self == nil {
			rpc.AddNode(conf.AppConfigInstance.MasterEndpoint)
		} else {
			logrus.Info("Node already exists")
			conf.AppConfigInstance.NodeID = self.UID
		}
	}
	router.Run(fmt.Sprintf("%v:%d", conf.AppConfigInstance.ListenAddr, conf.AppConfigInstance.APIPort))
}
