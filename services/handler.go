package services

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"time"
	"vorker/authz"
	"vorker/conf"
	"vorker/rpc"
	"vorker/services/agent"
	"vorker/services/appconf"
	"vorker/services/auth"
	"vorker/services/litefs"
	"vorker/services/node"
	proxyService "vorker/services/proxy"
	"vorker/services/workerd"
	"vorker/tunnel"
	"vorker/utils"
	"vorker/utils/database"

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
		if conf.IsMaster() {
			workerApi := api.Group("/worker", authz.JWTMiddleware())
			{
				workerApi.GET("/:uid", workerd.GetWorkerEndpoint)
				workerApi.GET("/flush/:uid", workerd.FlushEndpoint)
				workerApi.GET("/run/:uid", workerd.RunWorkerEndpoint)
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
			nodeAPI := api.Group("/node")
			{
				nodeAPI.GET("/all", authz.JWTMiddleware(), node.UserGetNodesEndpoint)
				nodeAPI.GET("/sync/:nodename", authz.JWTMiddleware(), node.SyncNodeEndpoint)
				nodeAPI.DELETE("/:nodename", authz.JWTMiddleware(), node.LeaveEndpoint)
			}
			api.GET("/allworkers", authz.JWTMiddleware(), workerd.GetAllWorkersEndpoint)
			api.GET("/vorker/config", appconf.GetEndpoint)
			api.POST("/auth/register", auth.RegisterEndpoint)
			api.POST("/auth/login", auth.LoginEndpoint)
			api.GET("/auth/logout", authz.JWTMiddleware(), auth.LogoutEndpoint)
		}
		agentAPI := api.Group("/agent")
		{
			if conf.IsMaster() {
				agentAPI.POST("/sync", authz.AgentAuthz(), workerd.AgentSyncWorkers)
				agentAPI.POST("/add", authz.AgentAuthz(), node.AddEndpoint)
				agentAPI.GET("/nodeinfo", authz.AgentAuthz(), node.GetNodeInfoEndpoint)
			} else {
				agentAPI.POST("/notify", authz.AgentAuthz(), agent.NotifyEndpoint)
			}
		}
	}

	proxy.Any("/*proxyPath", proxyService.Endpoint)
}

func InitTunnel() {
	if conf.IsMaster() {
		go tunnel.Serve()
		go tunnel.GetClient().Run(context.Background())
		go tunnel.InitSelfCliet()
	} else {
		go RegisterNodeToMaster()
		go tunnel.GetClient().Run(context.Background())
	}
	go litefs.InitTunnel()
	go litefs.RunService()
}

func Run(f embed.FS) {
	InitTunnel()
	go proxy.Run(fmt.Sprintf("%v:%d", conf.AppConfigInstance.ListenAddr, conf.AppConfigInstance.WorkerPort))
	go database.InitDB()

	if conf.IsMaster() {
		HandleStaticFile(f)
	} else {
		router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"code": 0, "msg": "ok"}) })
	}

	router.Run(fmt.Sprintf("%v:%d", conf.AppConfigInstance.ListenAddr, conf.AppConfigInstance.APIPort))
}

func HandleStaticFile(f embed.FS) {
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
	router.StaticFileFS("/nodes", "nodes.html", http.FS(fp))
	router.NoRoute(func(c *gin.Context) {
		c.FileFromFS(c.Request.URL.Path, http.FS(fp))
	})
}

func RegisterNodeToMaster() {
	if conf.IsMaster() {
		return
	}
	if conf.AppConfigInstance.LitefsEnabled {
		utils.WaitForPort("localhost", conf.AppConfigInstance.LitefsPrimaryPort)
	}
	for {
		logrus.Info("Registering node to master...")
		self, err := rpc.GetNode(conf.AppConfigInstance.MasterEndpoint)
		if err != nil || self == nil {
			err := rpc.AddNode(conf.AppConfigInstance.MasterEndpoint)
			if err != nil {
				logrus.WithError(err).Error("Add node failed.. retrying for 5 seconds")
				time.Sleep(5 * time.Second)
			} else {
				logrus.Info("Node added successfully")
			}
			continue
		} else {
			logrus.Info("Node already exists")
			conf.AppConfigInstance.NodeID = self.UID
		}
		tun, err := tunnel.GetClient().Query(conf.AppConfigInstance.NodeID)
		if err != nil || tun == nil {
			logrus.Warnf("Query tunnel failed, err: %v, try to add tunnel", err)
			tunnel.GetClient().Add(conf.AppConfigInstance.NodeID, utils.NodeHostPrefix(
				conf.AppConfigInstance.NodeName, conf.AppConfigInstance.NodeID),
				int(conf.AppConfigInstance.APIPort))
		} else {
			logrus.Info("Tunnel already exists, skip adding")
		}
		if conf.AppConfigInstance.EnableAutoSync {
			agent.SyncCall()
		}
		time.Sleep(30 * time.Second)
	}
}
