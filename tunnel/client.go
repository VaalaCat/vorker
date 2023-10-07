package tunnel

import (
	"fmt"
	"vorker/conf"
	"vorker/defs"
	"vorker/entities"
	"vorker/utils"

	"github.com/VaalaCat/tunnel/client"
	"github.com/sirupsen/logrus"
)

func InitAgent(allWorkers []*entities.Worker, allNodes []*entities.Node) {
	if conf.IsMaster() {
		return
	}
	for _, worker := range allWorkers {
		if worker.GetNodeName() == conf.AppConfigInstance.NodeName {
			continue
		}
		Add(worker.GetTunnelID(), worker.GetHostName(), worker.GetPort())
	}

	for _, node := range allNodes {
		Add(node.GetUID(), conf.AppConfigInstance.TunnelHost, int32(conf.AppConfigInstance.APIPort))
	}
}

func Add(clientID, forwardHostname string, forwardPort int32) {
	if conf.IsMaster() {
		return
	}
	freePort, err := utils.GetAvailablePort(defs.DefaultHostName)
	if err != nil {
		logrus.Errorf("tunnel add failed to get available port, err: %v", err)
	}
	client.RunClient(conf.AppConfigInstance.RPCHost, conf.AppConfigInstance.RPCPort, int64(freePort),
		fmt.Sprintf("%s:%d", forwardHostname, forwardPort), clientID,
	)
}

func Delete(clientID string) {
	if conf.IsMaster() {
		return
	}
	client.DeleteClient(conf.AppConfigInstance.RPCHost, conf.AppConfigInstance.RPCPort, clientID)
}
