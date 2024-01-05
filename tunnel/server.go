package tunnel

import (
	"context"
	"strings"
	"vorker/conf"

	"github.com/fatedier/frp/pkg/config"
	"github.com/fatedier/frp/server"
	"github.com/sirupsen/logrus"
)

func Serve() {
	cfg := config.GetDefaultServerConf()
	cfg.BindPort = int(conf.AppConfigInstance.TunnelAPIPort)
	cfg.VhostHTTPPort = int(conf.AppConfigInstance.TunnelEntryPort)
	cfg.SubDomainHost = strings.Trim(conf.AppConfigInstance.WorkerURLSuffix, ".")
	cfg.Token = conf.AppConfigInstance.TunnelToken
	cfg.TLSOnly = true
	svr, err := server.NewService(cfg)
	logrus.Infof("tunnel server listen on %v", conf.AppConfigInstance.TunnelAPIPort)
	if err != nil {
		return
	}
	svr.Run(context.Background())
}
