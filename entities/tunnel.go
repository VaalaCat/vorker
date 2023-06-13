package entities

import (
	sync "sync"

	"github.com/sirupsen/logrus"
)

type Tunnel struct {
	tunnelMap *sync.Map // key is worker's name, value is endpoint uuid
}

var (
	tunnelMap = &Tunnel{
		tunnelMap: &sync.Map{},
	}
)

func GetTunnel() *Tunnel {
	return tunnelMap
}

func (t *Tunnel) InitTunnelMap(
	workerList *WorkerList,
	nodesMap map[string]string,
) { // value of nodesMap is node's uuid, key is node's name
	if len(nodesMap) == 0 {
		logrus.Warn("no nodes found, please add nodes to enable tunnel service")
		return
	}

	for _, worker := range workerList.Workers {
		if nodeId, ok := nodesMap[worker.NodeName]; ok {
			t.tunnelMap.Store(worker.Name, nodeId)
		}
	}
}

func (t *Tunnel) GetTunnelMap() *sync.Map {
	return t.tunnelMap
}

func (t *Tunnel) AddTunnel(worker *Worker, nodeId string) {
	t.tunnelMap.Store(worker.Name, nodeId)
}

func (t *Tunnel) DeleteTunnel(worker *Worker) {
	t.tunnelMap.Delete(worker.Name)
}

func (t *Tunnel) GetAll() map[string]string {
	ans := make(map[string]string)
	t.tunnelMap.Range(func(key, value interface{}) bool {
		ans[key.(string)] = value.(string)
		return true
	})
	return ans
}
