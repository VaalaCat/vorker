package proxy

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"runtime/debug"
	"vorker/common"
	"vorker/conf"
	"vorker/entities"
	"vorker/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	proxy := entities.GetProxy()
	tunnel := entities.GetTunnel()
	workerRecords, err := models.AdminGetAllWorkers()
	if err != nil {
		logrus.Errorf("failed to get all workers, err: %v", err)
	}
	workerList := &entities.WorkerList{
		Workers: models.Trans2Entities(workerRecords),
	}

	if err != nil {
		logrus.Errorf("failed to get all nodes, err: %v", err)
	}

	proxy.InitProxyMap(workerList)
	tunnel.InitTunnelMap(workerList)
}

func Endpoint(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		}
	}()
	host := c.Request.Host
	c.Request.Host = host

	remote, err := url.Parse(fmt.Sprintf("http://%s:%d", conf.AppConfigInstance.TunnelHost,
		conf.AppConfigInstance.TunnelEntryPort))
	if err != nil {
		logrus.Panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(c.Writer, c.Request)
}
