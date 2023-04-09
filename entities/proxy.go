package entities

import sync "sync"

type Proxy struct {
	proxyMap *sync.Map
}

var (
	proxyMap = &Proxy{
		proxyMap: &sync.Map{},
	}
)

func GetProxy() *Proxy {
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
