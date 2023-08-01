package gost

import (
	"github.com/go-gost/x/config"
	"github.com/sirupsen/logrus"
)

func (g *GostClient) PostChain(chainConf *config.ChainConfig) error {
	url := g.APIBaseURL + "/config/chains"
	rawResp, err := g.reqCli().R().SetBody(chainConf).Post(url)
	if err != nil || rawResp.Response.StatusCode >= 299 {
		logrus.WithError(err).Errorf("post chain error, resp: %+v", rawResp)
		return err
	}
	logrus.Infof("post chain success, resp: %+v", rawResp)
	return nil
}

func (g *GostClient) DeleteChain(chainName string) error {
	url := g.APIBaseURL + "/config/chains/" + chainName
	rawResp, err := g.reqCli().R().Delete(url)
	if err != nil || rawResp.Response.StatusCode >= 299 {
		logrus.WithError(err).Errorf("delete chain error, resp: %+v", rawResp)
		return err
	}
	logrus.Infof("delete chain success, resp: %+v", rawResp)
	return nil
}

func (g *GostClient) PutChain(chainName string, chainConf *config.ChainConfig) error {
	url := g.APIBaseURL + "/config/chains/" + chainName
	rawResp, err := g.reqCli().R().SetBody(chainConf).Put(url)
	if err != nil || rawResp.Response.StatusCode >= 299 {
		logrus.WithError(err).Errorf("put chain error, resp: %+v", rawResp)
		return err
	}
	logrus.Infof("put chain success, resp: %+v", rawResp)
	return nil
}
