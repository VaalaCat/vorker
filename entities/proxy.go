package entities

import sync "sync"

type Proxy struct {
	proxyMap *sync.Map
}

var (
	proxyMap *Proxy
)

func GetProxy() *Proxy {
	if proxyMap == nil {
		proxyMap = &Proxy{
			proxyMap: &sync.Map{},
		}
	}
	return proxyMap
}

func (p *Proxy) InitProxyMap(workerList *WorkerList) {
	for _, worker := range workerList.Workers {
		p.proxyMap.Store(worker.Name, worker.Port)
	}
}

func (p *Proxy) GetProxyPort(name string) int32 {
	ans, ok := p.proxyMap.Load(name)
	if !ok {
		return 0
	}
	return ans.(int32)
}

func (p *Proxy) AddProxyPort(name string, port int32) {
	p.proxyMap.Store(name, port)
}

func (p *Proxy) DeleteProxyPort(name string) {
	p.proxyMap.Delete(name)
}

// GetAll returns a map, key is worker's name, value is worker's port
func (p *Proxy) GetAll() map[string]int32 {
	ans := make(map[string]int32)
	p.proxyMap.Range(func(key, value interface{}) bool {
		ans[key.(string)] = value.(int32)
		return true
	})
	return ans
}
