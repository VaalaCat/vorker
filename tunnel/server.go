package tunnel

import (
	"context"
	"strings"
	"time"
	"vorker/conf"
	"vorker/utils"

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

func InitSelfCliet() {
	utils.WaitForPort("localhost", conf.AppConfigInstance.LitefsPrimaryPort)
	for {
		if len(conf.AppConfigInstance.NodeID) == 0 {
			logger(context.Background(), "InitSelfCliet").Error("node is not initialized, retrying after 5 seconds")
			time.Sleep(5 * time.Second)
			continue
		}
		if err := GetClient().Add(conf.AppConfigInstance.NodeID, utils.NodeHostPrefix(
			conf.AppConfigInstance.NodeName, conf.AppConfigInstance.NodeID),
			int(conf.AppConfigInstance.APIPort)); err != nil {
			logger(context.Background(), "InitSelfCliet").Errorf("add tunnel failed, err: %v, retrying after 5 seconds", err)
			time.Sleep(5 * time.Second)
			continue
		}
		logger(context.Background(), "InitSelfCliet").Info("tunnel added successfully")
		break
	}
}
