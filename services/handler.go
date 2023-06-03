package services

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"voker/authz"
	"voker/conf"
	"voker/services/appconf"
	"voker/services/auth"
	proxyService "voker/services/proxy"
	"voker/services/workerd"
	"voker/utils"

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
	router.Use(utils.GinMiddleware("http://localhost:3000"))
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
	router.Run(fmt.Sprintf("%v:%d", conf.AppConfigInstance.ListenAddr, conf.AppConfigInstance.APIPort))
}
