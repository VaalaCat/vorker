package gost

import (
	"fmt"
	"vorker/conf"

	"github.com/go-gost/x/api"
	"github.com/go-gost/x/config"
	"github.com/imroc/req/v3"
	"github.com/sirupsen/logrus"
)

type GostInterface interface {
	reqCli() *req.Client
	PostConfig() error

	PostService(servConf *config.ServiceConfig) error
	DeleteService(serviceName string) error
	PutService(servicename string, servConf *config.ServiceConfig) error

	PostChain(chainConf *config.ChainConfig) error
	DeleteChain(chainName string) error
	PutChain(chainName string, chainConf *config.ChainConfig) error

	PostIngress(ingressConf *config.IngressConfig) error
	DeleteIngress(ingressName string) error
	PutIngress(ingressName string, ingressConf *config.IngressConfig) error
}

type GostClient struct {
	APIBaseURL string
	UserName   string
	Password   string
}

func (g *GostClient) reqCli() *req.Client {
	cli := req.C().
		SetCommonBasicAuth(conf.AppConfigInstance.TunnelUsername, conf.AppConfigInstance.TunnelPassword)
	return cli
}

func NewGostClient() GostInterface {
	return &GostClient{
		APIBaseURL: fmt.Sprintf("http://%s%s", conf.AppConfigInstance.TunnelHost, conf.AppConfigInstance.TunnelAPIPort),
		UserName:   conf.AppConfigInstance.TunnelUsername,
		Password:   conf.AppConfigInstance.TunnelPassword,
	}
}

func (g *GostClient) PostConfig() error {
	if conf.AppConfigInstance.RunMode != "master" {
		return nil
	}

	url := g.APIBaseURL + "/config"
	resp := &api.Response{}
	rawResp, err := g.reqCli().R().SetSuccessResult(resp).Post(url)
	if err != nil || rawResp.Response.StatusCode >= 299 {
		logrus.WithError(err).Errorf("post ingress error, resp: %+v", rawResp)
		return err
	}
	logrus.Infof("post ingress success, resp: %+v", resp)
	return nil
}
