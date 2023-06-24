package tunnel

import (
	"fmt"
	"runtime/debug"
	"sort"
	"strings"
	"vorker/common"
	"vorker/conf"
	"vorker/defs"
	"vorker/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetIngressConf(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		}
	}()
	allTunnel, err := models.AdminGetAllWorkersTunnelMap()
	if err != nil {
		logrus.Errorf("get all workers failed: %v", err)
		common.RespErr(c, defs.CodeInternalError, err.Error(), nil)
		return
	}

	workersRule := genWorkersRule(allTunnel)
	// get all nodes name from database
	allNodes, err := models.AdminGetAllNodesMap()
	if err != nil {
		logrus.Errorf("get all nodes map failed: %v", err)
		common.RespErr(c, defs.CodeInternalError, err.Error(), nil)
		return
	}
	nodesRule := genNodesRule(allNodes)
	conf := buildIngressConf(append(workersRule, nodesRule...))
	c.String(200, conf)
}

func buildIngressConf(rules []string) string {
	ans := strings.Join(rules, "\n")
	return ans
}

func genWorkersRule(tunnels map[string]string) []string {
	rules := []string{}
	for workerName, tunnelID := range tunnels {
		rules = append(rules, fmt.Sprintf("%s%s %s", workerName,
			conf.AppConfigInstance.WorkerURLSuffix, tunnelID))
	}
	// sort rules
	sort.Strings(rules)
	return rules
}

func genNodesRule(nodes map[string]string) []string {
	rules := []string{}
	for nodeName, nodeId := range nodes {
		rules = append(rules, fmt.Sprintf("%s%s %s", nodeName,
			nodeId, nodeId))
	}
	sort.Strings(rules)
	return rules
}
