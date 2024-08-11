package proxy

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"runtime/debug"
	"vorker/common"
	"vorker/conf"
	"vorker/models"
	"vorker/tunnel"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Endpoint(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		}
	}()
	host := c.Request.Host
	c.Request.Host = host
	workerName := host[:len(host)-len(conf.AppConfigInstance.WorkerURLSuffix)]
	worker, err := models.AdminGetWorkerByName(workerName)
	if err != nil {
		logrus.Errorf("failed to get worker by name, err: %v", err)
		common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		return
	}

	var remote *url.URL
	if worker.GetNodeName() == conf.AppConfigInstance.NodeName {
		workerPort, ok := tunnel.GetPortManager().GetWorkerPort(c, worker.GetUID())
		if !ok {
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
			return
		}
		remote, err = url.Parse(fmt.Sprintf("http://%s:%d", worker.GetHostName(), workerPort))
		if err != nil {
			logrus.Panic(err)
		}
	} else {
		remote, err = url.Parse(fmt.Sprintf("http://%s:%d",
			conf.AppConfigInstance.TunnelHost, conf.AppConfigInstance.TunnelEntryPort))
		if err != nil {
			logrus.Panic(err)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(c.Writer, c.Request)
}
