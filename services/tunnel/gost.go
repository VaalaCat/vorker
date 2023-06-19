package tunnel

import (
	"fmt"
	"sort"
	"strings"
	"voker/common"
	"voker/conf"
	"voker/defs"
	"voker/entities"
	"voker/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetIngressConf(c *gin.Context) {
	allTunnel := entities.GetTunnel().GetAll()
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
