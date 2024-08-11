package tunnel

import (
	"context"
	"fmt"
	"time"
	"vorker/conf"
	"vorker/utils"

	"github.com/fatedier/frp/client"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

type ClientHandler interface {
	Run(ctx context.Context)
	Add(clientID, routeHostname string, forwardPort int) error
	AddService(serviceName string, servicePort int) error
	AddVisitor(servicename string, lcoalPort int) error
	Delete(clientID string) error
	Query(clientID string) (v1.ProxyConfigurer, error)
}

type Client struct {
	proxyConf   *utils.SyncMap[string, v1.ProxyConfigurer]
	visitorConf *utils.SyncMap[string, v1.VisitorConfigurer]
	common      *v1.ClientCommonConfig
	cli         *client.Service
}

var (
	cli ClientHandler
)

func NewClientHandler() *Client {
	cfg := v1.ClientCommonConfig{
		Auth: v1.AuthClientConfig{
			Method: "token",
			Token:  conf.AppConfigInstance.TunnelToken,
		},
		ServerAddr:    conf.AppConfigInstance.TunnelHost,
		ServerPort:    int(conf.AppConfigInstance.TunnelAPIPort),
		LoginFailExit: lo.ToPtr(false),
	}

	c, err := client.NewService(client.ServiceOptions{
		Common: &cfg,
	})
	if err != nil {
		logrus.WithError(err).Error("New client failed")
		return nil
	}

	return &Client{
		proxyConf:   &utils.SyncMap[string, v1.ProxyConfigurer]{},
		visitorConf: &utils.SyncMap[string, v1.VisitorConfigurer]{},
		cli:         c,
		common:      &cfg,
	}
}

func GetClient() ClientHandler {
	if cli == nil {
		cli = NewClientHandler()
	}
	return cli
}

// Add implements ClientHandler.
func (c *Client) Add(clientID, routeHostname string, forwardPort int) error {
	var newCfg v1.ProxyConfigurer = &v1.HTTPProxyConfig{
		ProxyBaseConfig: v1.ProxyBaseConfig{
			Name: clientID,
			Type: "http",
			ProxyBackend: v1.ProxyBackend{
				LocalIP:   "127.0.0.1",
				LocalPort: forwardPort,
			},
		},
		DomainConfig: v1.DomainConfig{
			SubDomain: routeHostname,
		},
	}
	newCfg.Complete("")
	if _, ok := c.proxyConf.LoadOrStore(clientID, newCfg); ok {
		logger(context.Background(), "Client.Add").Errorf("client %s already exists", clientID)
		return nil
	}

	err := c.cli.UpdateAllConfigurer(lo.Values(c.proxyConf.ToMap()), lo.Values(c.visitorConf.ToMap()))
	if err != nil {
		logger(context.Background(), "Client.Add").WithError(err).
			Errorf("reload conf failed, config is: %+v", c.proxyConf.ToMap())
		return err
	}
	logger(context.Background(), "Client.Add").Infof("client %s added successfully", clientID)
	return nil
}

// AddService implements ClientHandler.
func (c *Client) AddService(serviceName string, servicePort int) error {
	var newCfg v1.ProxyConfigurer = &v1.STCPProxyConfig{
		ProxyBaseConfig: v1.ProxyBaseConfig{
			Name: serviceName,
			Type: "stcp",
			Transport: v1.ProxyTransport{
				UseEncryption: true,
			},
			ProxyBackend: v1.ProxyBackend{
				LocalIP:   "127.0.0.1",
				LocalPort: servicePort,
			},
		},
		Secretkey: conf.AppConfigInstance.TunnelToken,
	}
	newCfg.Complete("")

	if _, ok := c.proxyConf.LoadOrStore(serviceName, newCfg); ok {
		logger(context.Background(), "Client.AddService").Errorf("service %s already exists", serviceName)
		return nil
	}

	err := c.cli.UpdateAllConfigurer(lo.Values(c.proxyConf.ToMap()), lo.Values(c.visitorConf.ToMap()))
	if err != nil {
		logger(context.Background(), "Client.AddService").WithError(err).
			Errorf("reload conf failed, config is: %+v", c.proxyConf.ToMap())
		return err
	}
	logger(context.Background(), "Client.AddService").Infof("service %s added successfully", serviceName)
	return nil
}

// AddVisitor implements ClientHandler.
func (c *Client) AddVisitor(serviceName string, lcoalPort int) error {
	var newCfg v1.VisitorConfigurer = &v1.STCPVisitorConfig{
		VisitorBaseConfig: v1.VisitorBaseConfig{
			Name: fmt.Sprintf("%s-visitor", serviceName),
			Type: "stcp",
			Transport: v1.VisitorTransport{
				UseEncryption: true,
			},
			BindAddr:   "127.0.0.1",
			BindPort:   lcoalPort,
			ServerName: serviceName,
			SecretKey:  conf.AppConfigInstance.TunnelToken,
		},
	}
	newCfg.Complete(c.common)

	if _, ok := c.visitorConf.LoadOrStore(serviceName, newCfg); ok {
		logger(context.Background(), "Client.AddVisitor").Errorf("visitor for serivce %s already exists", serviceName)
		return nil
	}

	err := c.cli.UpdateAllConfigurer(lo.Values(c.proxyConf.ToMap()), lo.Values(c.visitorConf.ToMap()))
	if err != nil {
		logger(context.Background(), "Client.AddVisitor").WithError(err).
			Errorf("reload conf failed, config is: %+v", c.visitorConf.ToMap())
		return err
	}
	logger(context.Background(), "Client.AddVisitor").Infof("visitor for service %s added successfully", serviceName)
	return nil
}

// Delete implements ClientHandler.
func (c *Client) Delete(clientID string) error {
	if _, ok := c.proxyConf.Load(clientID); !ok {
		logger(context.Background(), "Client.Delete").Errorf("client %s not exists", clientID)
		return nil
	}

	c.proxyConf.Delete(clientID)
	err := c.cli.UpdateAllConfigurer(lo.Values(c.proxyConf.ToMap()), lo.Values(c.visitorConf.ToMap()))
	if err != nil {
		logger(context.Background(), "Client.Delete").WithError(err).
			Errorf("reload conf failed, config is: %+v", c.proxyConf.ToMap())
		return err
	}
	logger(context.Background(), "Client.Delete").Infof("client %s deleted successfully", clientID)
	return nil
}

// Query implements ClientHandler.
func (c *Client) Query(clientID string) (v1.ProxyConfigurer, error) {
	if proxyConf, ok := c.proxyConf.Load(clientID); ok {
		return proxyConf, nil
	}
	logger(context.Background(), "Client.Query").Errorf("client %s not exists", clientID)
	return nil, fmt.Errorf("client %s not exists", clientID)
}

// Run implements ClientHandler.
func (c *Client) Run(ctx context.Context) {
	for {
		if err := c.cli.Run(ctx); err != nil {
			logger(ctx, "Client.Run").WithError(err).Error("client run failed, retrying for 5 seconds")
			time.Sleep(5 * time.Second)
		}
	}
}
