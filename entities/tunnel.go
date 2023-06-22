package entities

import (
	sync "sync"
)

type Tunnel struct {
	tunnelMap *sync.Map // key is worker's name, value is tunnel uuid
}

var (
	tunnelMap *Tunnel
)

func GetTunnel() *Tunnel {
	if tunnelMap == nil {
		tunnelMap = &Tunnel{
			tunnelMap: &sync.Map{},
		}
	}
	return tunnelMap
}

func (t *Tunnel) InitTunnelMap(
	workerList *WorkerList,
) {
	for _, worker := range workerList.Workers {
		t.tunnelMap.Store(worker.Name, worker.TunnelID)
	}
}

func (t *Tunnel) GetTunnelMap() *sync.Map {
	return t.tunnelMap
}

func (t *Tunnel) AddTunnel(worker *Worker) {
	t.tunnelMap.Store(worker.Name, worker.TunnelID)
}

func (t *Tunnel) DeleteTunnel(worker *Worker) {
	t.tunnelMap.Delete(worker.Name)
}

// GetAll returns a map, key is worker's name, value is tunnel uuid
func (t *Tunnel) GetAll() map[string]string {
	ans := make(map[string]string)
	t.tunnelMap.Range(func(key, value interface{}) bool {
		ans[key.(string)] = value.(string)
		return true
	})
	return ans
}

func CheckTunnel(workerName string) bool {
	_, ok := tunnelMap.tunnelMap.Load(workerName)
	return ok
}
