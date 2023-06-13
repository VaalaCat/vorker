package proxy

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"strings"
	"voker/conf"
	"voker/entities"
	"voker/models"

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

	nodesMap, err := models.AdminGetAllNodesMap()
	if err != nil {
		logrus.Errorf("failed to get all nodes, err: %v", err)
	}

	proxy.InitProxyMap(workerList)
	tunnel.InitTunnelMap(workerList, nodesMap)
}

func Endpoint(c *gin.Context) {
	host := c.Request.Host
	name := strings.Split(host, ".")[0]
	port := entities.GetProxy().GetProxyPort(name)
	c.Request.Host = name
	if port == 0 {
		tunnel, err := url.Parse(fmt.Sprintf("http://localhost:%v", conf.AppConfigInstance.TunnelPort))
		if err != nil {
			logrus.Panic(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(tunnel)
		proxy.ServeHTTP(c.Writer, c.Request)
		return
	}

	remote, err := url.Parse(fmt.Sprintf("http://localhost:%v", port))
	if err != nil {
		logrus.Panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(c.Writer, c.Request)
}
