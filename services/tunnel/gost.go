package tunnel

import (
	"fmt"
	"sort"
	"strings"
	"voker/entities"

	"github.com/gin-gonic/gin"
)

func GetIngressConf(c *gin.Context) {
	c.String(200, genConf())
}

func genConf() string {
	allTunnel := entities.GetTunnel().GetAll()
	rules := []string{}
	for workerName, nodeId := range allTunnel {
		rules = append(rules, fmt.Sprintf("%s %s", workerName, nodeId))
	}
	// sort rules
	sort.Strings(rules)
	ans := strings.Join(rules, "\n")
	return ans
}
