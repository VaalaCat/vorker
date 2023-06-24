package models

import (
	"fmt"
	"vorker/conf"
	"vorker/entities"
)

func buildGostArgs(scheme, ip string, port int32, rhost string, tunnelID string) stringList {
	ans := make(stringList, 0)
	ans.Set("-L")
	ans.Set(fmt.Sprintf("rtcp://:0/%s:%d", ip, port))
	ans.Set("-F")
	ans.Set(fmt.Sprintf("%s://%s:%s@%s?tunnel.id=%s", scheme,
		conf.AppConfigInstance.TunnelUsername,
		conf.AppConfigInstance.TunnelPassword,
		rhost, tunnelID))
	return ans
}

func buildGostPool(workers []*entities.Worker) map[string]stringList {
	ans := make(map[string]stringList, 0)
	for _, worker := range workers {
		ip, port := "127.0.0.1", worker.Port
		ans[worker.Name] =
			buildGostArgs(conf.AppConfigInstance.TunnelScheme,
				ip, port, conf.AppConfigInstance.TunnelRelayEndpoint,
				worker.TunnelID)
	}
	return ans
}
