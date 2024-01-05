package tunnel

import (
	"context"
	"fmt"
	"time"
	"vorker/conf"

	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/pkg/config"
	"github.com/sirupsen/logrus"
)

type ClientHandler interface {
	Run(ctx context.Context)
	Add(clientID, routeHostname string, forwardPort int) error
	AddService(serviceName string, servicePort int) error
	AddVisitor(servicename string, lcoalPort int) error
	Delete(clientID string) error
	Query(clientID string) (config.ProxyConf, error)
}

type Client struct {
	proxyConf   map[string]config.ProxyConf
	visitorConf map[string]config.VisitorConf
	cli         *client.Service
}

var (
	cli ClientHandler
)

func NewClientHandler() *Client {
	cfg := config.GetDefaultClientConf()
	cfg.ServerAddr = conf.AppConfigInstance.TunnelHost
	cfg.ServerPort = int(conf.AppConfigInstance.TunnelAPIPort)
	cfg.Token = conf.AppConfigInstance.TunnelToken
	proxyConf := map[string]config.ProxyConf{}
	visitorConf := map[string]config.VisitorConf{}
	c, err := client.NewService(cfg,
		proxyConf, visitorConf, "")
	if err != nil {
		logrus.WithError(err).Error("New client failed")
		return nil
	}
	return &Client{
		proxyConf:   proxyConf,
		visitorConf: visitorConf,
		cli:         c,
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
	newProxyConf := c.proxyConf
	if _, ok := newProxyConf[clientID]; ok {
		logger(context.Background(), "Client.Add").Errorf("client %s already exists", clientID)
		return nil
	}

	newProxyConf[clientID] = &config.HTTPProxyConf{
		BaseProxyConf: config.BaseProxyConf{
			ProxyName: clientID,
			ProxyType: "http",
			LocalSvrConf: config.LocalSvrConf{
				LocalIP:   "127.0.0.1",
				LocalPort: forwardPort,
			},
		},
		DomainConf: config.DomainConf{SubDomain: routeHostname},
	}
	c.proxyConf = newProxyConf

	return c.cli.ReloadConf(newProxyConf, c.visitorConf)
}

// AddService implements ClientHandler.
func (c *Client) AddService(serviceName string, servicePort int) error {
	newSerivceConf := c.proxyConf
	if _, ok := newSerivceConf[serviceName]; ok {
		logger(context.Background(), "Client.AddService").Errorf("service %s already exists", serviceName)
		return nil
	}

	newSerivceConf[serviceName] = &config.STCPProxyConf{
		BaseProxyConf: config.BaseProxyConf{
			ProxyName:     serviceName,
			ProxyType:     "stcp",
			UseEncryption: true,
			LocalSvrConf: config.LocalSvrConf{
				LocalIP:   "127.0.0.1",
				LocalPort: servicePort,
			},
		},
		RoleServerCommonConf: config.RoleServerCommonConf{
			Sk: conf.AppConfigInstance.TunnelToken,
		},
	}

	c.proxyConf = newSerivceConf
	return c.cli.ReloadConf(c.proxyConf, c.visitorConf)
}

// AddVisitor implements ClientHandler.
func (c *Client) AddVisitor(serviceName string, lcoalPort int) error {
	newVisitorConf := c.visitorConf
	if _, ok := newVisitorConf[serviceName]; ok {
		logger(context.Background(), "Client.AddVisitor").Errorf("visitor for serivce %s already exists", serviceName)
		return nil
	}

	newVisitorConf[serviceName] = &config.STCPVisitorConf{
		BaseVisitorConf: config.BaseVisitorConf{
			ProxyName:     fmt.Sprintf("%s-visitor", serviceName),
			ProxyType:     "stcp",
			UseEncryption: true,
			BindAddr:      "127.0.0.1",
			BindPort:      lcoalPort,
			ServerName:    serviceName,
			Sk:            conf.AppConfigInstance.TunnelToken,
		},
	}

	c.visitorConf = newVisitorConf
	return c.cli.ReloadConf(c.proxyConf, c.visitorConf)
}

// Delete implements ClientHandler.
func (c *Client) Delete(clientID string) error {
	newProxyConf := c.proxyConf
	if _, ok := newProxyConf[clientID]; !ok {
		logger(context.Background(), "Client.Delete").Errorf("client %s not exists", clientID)
		return nil
	}

	delete(newProxyConf, clientID)
	c.proxyConf = newProxyConf
	return c.cli.ReloadConf(c.proxyConf, c.visitorConf)
}

// Query implements ClientHandler.
func (c *Client) Query(clientID string) (config.ProxyConf, error) {
	if proxyConf, ok := c.proxyConf[clientID]; ok {
		return proxyConf, nil
	}
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
