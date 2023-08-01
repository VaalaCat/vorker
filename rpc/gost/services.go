package gost

import (
	"fmt"

	"github.com/go-gost/x/api"
	gostConfig "github.com/go-gost/x/config"
	"github.com/sirupsen/logrus"
)

func (g *GostClient) PostService(servConf *gostConfig.ServiceConfig) error {
	url := g.APIBaseURL + "/config/services"
	resp := &api.Response{}
	rawResp, err := g.reqCli().R().SetBody(servConf).SetSuccessResult(resp).Post(url)
	if err != nil || rawResp.Response.StatusCode >= 299 {
		logrus.WithError(err).Errorf("post service error, resp: %+v", rawResp)
		return err
	}
	logrus.Infof("post service success, resp: %+v", resp)
	return nil
}

func (g *GostClient) DeleteService(serviceName string) error {
	url := g.APIBaseURL + fmt.Sprintf("/config/services/%s", serviceName)
	resp := &api.Response{}
	rawResp, err := g.reqCli().R().SetSuccessResult(resp).Delete(url)
	if err != nil || rawResp.Response.StatusCode >= 299 {
		logrus.WithError(err).Errorf("delete service error, resp: %+v", rawResp)
		return err
	}
	logrus.Infof("delete service success, resp: %+v", resp)
	return nil
}

func (g *GostClient) PutService(servicename string, servConf *gostConfig.ServiceConfig) error {
	url := g.APIBaseURL + fmt.Sprintf("/config/services/%s", servicename)
	resp := &api.Response{}
	rawResp, err := g.reqCli().R().SetBody(servConf).SetSuccessResult(resp).Put(url)
	if err != nil || rawResp.Response.StatusCode >= 299 {
		logrus.WithError(err).Errorf("put service error, resp: %+v", rawResp)
		return err
	}
	logrus.Infof("put service success, resp: %+v", resp)
	return nil
}
