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

func InitAgent(nodeWorkers []*entities.Worker, allWorkers map[string]string, allNodes map[string]string) {

}

func Add(clientID, forwardHostname string, forwardPort int32) {
	freePort, err := utils.GetAvailablePort(defs.DefaultHostName)
	if err != nil {
		logrus.Errorf("tunnel add failed to get available port, err: %v", err)
	}
	client.RunClient(conf.AppConfigInstance.RPCPort, int64(freePort),
		fmt.Sprintf("%s:%d", forwardHostname, forwardPort), clientID,
	)
}

func Delete(clientID string) {
	client.DeleteClient(conf.AppConfigInstance.RPCPort, clientID)
}
