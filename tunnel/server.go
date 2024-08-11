package tunnel

import (
	"context"
	"strings"
	"time"
	"vorker/conf"
	"vorker/utils"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/server"
	"github.com/sirupsen/logrus"
)

func Serve() {
	cfg := v1.ServerConfig{
		BindPort:      int(conf.AppConfigInstance.TunnelAPIPort),
		VhostHTTPPort: int(conf.AppConfigInstance.TunnelEntryPort),
		SubDomainHost: strings.Trim(conf.AppConfigInstance.WorkerURLSuffix, "."),
		Transport: v1.ServerTransportConfig{
			TLS: v1.TLSServerConfig{Force: true},
		},
		Auth: v1.AuthServerConfig{
			Method: "token", Token: conf.AppConfigInstance.TunnelToken,
		},
	}
	cfg.Complete()

	svr, err := server.NewService(&cfg)
	if err != nil {
		logrus.WithError(err).Error("new tunnel listen failed")
		return
	}
	logrus.Infof("tunnel server listen on %v", conf.AppConfigInstance.TunnelAPIPort)
	svr.Run(context.Background())
}

func InitSelfCliet() {
	if conf.AppConfigInstance.LitefsEnabled {
		utils.WaitForPort("localhost", conf.AppConfigInstance.LitefsPrimaryPort)
	}
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
