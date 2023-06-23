package proxy

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"runtime/debug"
	"vorker/common"
	"vorker/conf"

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

	remote, err := url.Parse(fmt.Sprintf("http://%s:%d", conf.AppConfigInstance.TunnelHost,
		conf.AppConfigInstance.TunnelEntryPort))
	if err != nil {
		logrus.Panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(c.Writer, c.Request)
}
