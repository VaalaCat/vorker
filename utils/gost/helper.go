package gost

import (
	"fmt"
	"voker/conf"
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

func buildGostPool(tunnelMap map[string]string, workerMap map[string]int32) map[string]stringList {
	ans := make(map[string]stringList, 0)
	for workerName, tunnelID := range tunnelMap {
		ip, port := "127.0.0.1", workerMap[workerName]
		ans[workerName] =
			buildGostArgs(conf.AppConfigInstance.TunnelScheme,
				ip, port, conf.AppConfigInstance.TunnelRelayEndpoint,
				tunnelID)
	}
	return ans
}
