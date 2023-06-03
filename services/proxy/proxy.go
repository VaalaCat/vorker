package proxy

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"strings"
	"voker/entities"
	"voker/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	proxy := entities.GetProxy()
	workerRecords, err := models.AdminGetAllWorkers()
	if err != nil {
		logrus.Errorf("failed to get all workers, err: %v", err)
	}
	workerList := &entities.WorkerList{
		Workers: models.Trans2Entities(workerRecords),
	}
	proxy.InitProxyMap(workerList)
}

func Endpoint(c *gin.Context) {
	host := c.Request.Host
	name := strings.Split(host, ".")[0]
	port := entities.GetProxy().GetProxyPort(name)
	if port == 0 {
		c.JSON(404, gin.H{"code": 1, "error": "not found"})
		return
	}

	remote, err := url.Parse(fmt.Sprintf("http://localhost:%v", port))
	if err != nil {
		logrus.Panic(err)
	}

	c.Request.URL.Path = c.Copy().Param("name")
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(c.Writer, c.Request)
}
