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
	Delete(clientID string) error
	Query(clientID string) (config.ProxyConf, error)
}

type Client struct {
	proxyConf map[string]config.ProxyConf
	cli       *client.Service
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
	c, err := client.NewService(cfg,
		proxyConf, nil, "")
	if err != nil {
		logrus.WithError(err).Error("New client failed")
		return nil
	}
	return &Client{
		proxyConf: proxyConf,
		cli:       c,
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
				LocalPort: int(forwardPort),
			},
		},
		DomainConf: config.DomainConf{SubDomain: routeHostname},
	}

	return c.cli.ReloadConf(newProxyConf, nil)
}

// Delete implements ClientHandler.
func (c *Client) Delete(clientID string) error {
	newProxyConf := c.proxyConf
	if _, ok := newProxyConf[clientID]; !ok {
		logger(context.Background(), "Client.Delete").Errorf("client %s not exists", clientID)
		return nil
	}

	delete(newProxyConf, clientID)
	return c.cli.ReloadConf(newProxyConf, nil)
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
	go func() {
		for {
			if err := c.cli.Run(ctx); err != nil {
				logger(ctx, "Client.Run").WithError(err).Error("client run failed, retrying for 5 seconds")
				time.Sleep(5 * time.Second)
			}
		}
	}()
}
