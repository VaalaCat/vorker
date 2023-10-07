package proxy

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"runtime/debug"
	"vorker/common"
	"vorker/conf"
	"vorker/models"

	"github.com/VaalaCat/tunnel/forwarder"
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
		remote, err = url.Parse(fmt.Sprintf("http://%s:%d", worker.GetHostName(), worker.GetPort()))
		if err != nil {
			logrus.Panic(err)
		}
	} else {
		tun, err := forwarder.GetListener().GetTunnelInfo(worker.GetTunnelID())
		if err != nil {
			logrus.Errorf("failed to get tunnel info, err: %v", err)
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
			return
		}

		remote, err = url.Parse(fmt.Sprintf("http://%s:%d", conf.AppConfigInstance.TunnelHost,
			tun.GetPort()))
		if err != nil {
			logrus.Panic(err)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(c.Writer, c.Request)
}
