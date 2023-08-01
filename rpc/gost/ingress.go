package gost

import (
	"github.com/go-gost/x/api"
	gostConfig "github.com/go-gost/x/config"
	"github.com/sirupsen/logrus"
)

func (g *GostClient) PostIngress(ingressConf *gostConfig.IngressConfig) error {
	url := g.APIBaseURL + "/config/ingresses"
	resp := &api.Response{}
	rawResp, err := g.reqCli().R().SetBody(ingressConf).SetSuccessResult(resp).Post(url)
	if err != nil || rawResp.Response.StatusCode >= 299 {
		logrus.WithError(err).Errorf("post ingress error, resp: %+v", rawResp)
		return err
	}
	logrus.Infof("post ingress success, resp: %+v", resp)
	return nil
}

func (g *GostClient) DeleteIngress(ingressName string) error {
	url := g.APIBaseURL + "/config/ingresses/" + ingressName
	resp := &api.Response{}
	rawResp, err := g.reqCli().R().SetSuccessResult(resp).Delete(url)
	if err != nil || rawResp.Response.StatusCode >= 299 {
		logrus.WithError(err).Errorf("delete ingress error, resp: %+v", rawResp)
		return err
	}
	logrus.Infof("delete ingress success, resp: %+v", resp)
	return nil
}

func (g *GostClient) PutIngress(ingressName string, ingressConf *gostConfig.IngressConfig) error {
	url := g.APIBaseURL + "/config/ingresses/" + ingressName
	resp := &api.Response{}
	rawResp, err := g.reqCli().R().SetBody(ingressConf).SetSuccessResult(resp).Put(url)
	if err != nil || rawResp.Response.StatusCode >= 299 {
		logrus.WithError(err).Errorf("put ingress error, resp: %+v", rawResp)
		return err
	}
	logrus.Infof("put ingress success, resp: %+v", resp)
	return nil
}
