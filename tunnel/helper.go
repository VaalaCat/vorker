package tunnel

import (
	"fmt"
	"runtime/debug"
	"sort"
	"vorker/conf"
	"vorker/entities"

	"github.com/go-gost/x/config"
	"github.com/sirupsen/logrus"
)

const (
	VorkerDefaultChainName   = "vorker-chain-default"
	VorkerDefaultHopName     = "vorker-hop-default"
	VorkerDefaultNodeName    = "vorker-node-default"
	VorkerDefaultServiceName = "vorker-service-default"
	VorkerDefaultIngressName = "vorker-ingress-default"
)

const (
	TypeRelay = "relay"
	TypeWSS   = "wss"
	TypeWS    = "ws"
	TypeRTCP  = "rtcp"
)

const (
	MetadataKeyTunnelID = "tunnel.id"
	MetadataKeySniffing = "sniffing"
)

func chainConfigTemplate() *config.ChainConfig {
	return &config.ChainConfig{
		Name: VorkerDefaultChainName,
		Hops: []*config.HopConfig{
			{
				Name: VorkerDefaultHopName,
				Nodes: []*config.NodeConfig{
					{Name: VorkerDefaultNodeName, Addr: conf.AppConfigInstance.TunnelRelayEndpoint,
						Connector: &config.ConnectorConfig{
							Type:     TypeRelay,
							Metadata: map[string]interface{}{MetadataKeyTunnelID: conf.AppConfigInstance.NodeID},
							Auth: &config.AuthConfig{
								Username: conf.AppConfigInstance.TunnelUsername,
								Password: conf.AppConfigInstance.TunnelPassword,
							},
						},
						Dialer: &config.DialerConfig{Type: TypeWS},
					},
				},
			},
		},
	}
}

func serviceConfigTemplate(nodes []*config.ForwardNodeConfig) *config.ServiceConfig {
	return &config.ServiceConfig{
		Name: VorkerDefaultServiceName,
		Addr: ":0",
		Handler: &config.HandlerConfig{
			Type:     TypeRTCP,
			Metadata: map[string]interface{}{MetadataKeySniffing: true},
		},
		Listener: &config.ListenerConfig{
			Type:  TypeRTCP,
			Chain: VorkerDefaultChainName,
		},
		Forwarder: &config.ForwarderConfig{Nodes: nodes},
	}
}

func forwardNodesTemplate(workers []*entities.Worker) []*config.ForwardNodeConfig {
	resp := []*config.ForwardNodeConfig{}
	for _, worker := range workers {
		tmp := &config.ForwardNodeConfig{
			Host: fmt.Sprintf("%s%s", worker.Name, conf.AppConfigInstance.WorkerURLSuffix),
			Addr: fmt.Sprintf("%s:%d", worker.HostName, worker.Port),
			Name: worker.Name,
		}
		resp = append(resp, tmp)
	}
	return resp
}

func newForwardNode(tunnelID string, tunnelName string, tunnelPort int32) *config.ForwardNodeConfig {
	return &config.ForwardNodeConfig{
		Host: fmt.Sprintf("%s%s", tunnelName, conf.AppConfigInstance.WorkerURLSuffix),
		Addr: fmt.Sprintf("%s:%d", conf.AppConfigInstance.DefaultWorkerHost, tunnelPort),
		Name: tunnelName,
	}
}

func ingressConfigTemplate(rules []*config.IngressRuleConfig) *config.IngressConfig {
	return &config.IngressConfig{
		Name:  VorkerDefaultIngressName,
		Rules: rules,
	}
}

func genIngressRules(allTunnel, allNodes map[string]string) []*config.IngressRuleConfig {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
		}
	}()
	workersRule := genWorkersRule(allTunnel)
	sort.Slice(workersRule, func(i, j int) bool {
		return workersRule[i].Hostname < workersRule[j].Hostname
	})
	// get all nodes name from database
	nodesRule := genNodesRule(allNodes)
	sort.Slice(nodesRule, func(i, j int) bool {
		return nodesRule[i].Hostname < nodesRule[j].Hostname
	})

	return append(workersRule, nodesRule...)
}

func genWorkersRule(tunnels map[string]string) []*config.IngressRuleConfig {
	rules := []*config.IngressRuleConfig{}
	for workerName, tunnelID := range tunnels {
		hostname := fmt.Sprintf("%s%s", workerName, conf.AppConfigInstance.WorkerURLSuffix)
		rules = append(rules, &config.IngressRuleConfig{
			Hostname: hostname,
			Endpoint: tunnelID,
		})
	}
	return rules
}

func genNodesRule(nodes map[string]string) []*config.IngressRuleConfig {
	rules := []*config.IngressRuleConfig{}
	for nodeName, nodeId := range nodes {
		hostname := fmt.Sprintf("%s%s%s", nodeName, nodeId, conf.AppConfigInstance.WorkerURLSuffix)
		rules = append(rules, &config.IngressRuleConfig{
			Hostname: hostname,
			Endpoint: nodeId,
		})
	}
	return rules
}
