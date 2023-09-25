package tunnel

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"runtime/debug"
	"time"
	"vorker/conf"
	"vorker/entities"
	"vorker/rpc/gost"

	"github.com/go-gost/x/config"
	"github.com/judwhite/go-svc"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

func init() {

}

func InitTunnelAgent(workers []*entities.Worker, allWorkers, allNodes map[string]string) {
	// TODO: init gost client call gost api to add default chain
	for {
		if err := gost.NewGostClient().PostChain(chainConfigTemplate()); err != nil {
			logrus.WithError(err).Errorf("init tunnel agent post chain error")
			time.Sleep(3 * time.Second)
		} else {
			logrus.Infof("init tunnel agent post chain success")
			break
		}
	}
	for {
		if err := gost.NewGostClient().PostService(serviceConfigTemplate(forwardNodesTemplate(workers))); err != nil {
			logrus.WithError(err).Errorf("init tunnel agent post service error")
			time.Sleep(3 * time.Second)
		} else {
			logrus.Infof("init tunnel agent post service success")
			break
		}
	}
	rules := genIngressRules(allWorkers, allNodes)
	for {
		if conf.AppConfigInstance.RunMode == "agent" {
			break
		}
		if err := gost.NewGostClient().PostIngress(ingressConfigTemplate(rules)); err != nil {
			logrus.WithError(err).Errorf("init tunnel agent post ingress error")
			time.Sleep(3 * time.Second)
		} else {
			logrus.Infof("init tunnel agent post ingress success")
			break
		}
	}
	gost.NewGostClient().PostConfig()
}

func Add(tunnelID string, tunnelName string,
	tunnelPort int32, workers []*entities.Worker,
	allWorkers, allNodes map[string]string) error {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
		}
	}()

	oldNodes := forwardNodesTemplate(workers)
	newNode := newForwardNode(tunnelID, tunnelName, tunnelPort)
	newNodes := append(oldNodes, newNode)
	newServConf := serviceConfigTemplate(newNodes)

	maxRetry := 10
	var err error
	for i := 0; i < maxRetry; i++ {
		if err = gost.NewGostClient().PutService(VorkerDefaultServiceName, newServConf); err != nil {
			logrus.WithError(err).Errorf("add tunnel error")
			time.Sleep(3 * time.Second)
			continue
		} else {
			logrus.Infof("add tunnel success")
		}
	}
	if conf.AppConfigInstance.RunMode == "agent" {
		return err
	}

	rules := genIngressRules(allWorkers, allNodes)
	newHostname := fmt.Sprintf("%s%s", tunnelName,
		conf.AppConfigInstance.WorkerURLSuffix)
	newRules := append(rules, &config.IngressRuleConfig{
		Hostname: newHostname})
	ingressConf := ingressConfigTemplate(newRules)
	for i := 0; i < maxRetry; i++ {
		if err = gost.NewGostClient().PutIngress(VorkerDefaultIngressName, ingressConf); err != nil {
			logrus.WithError(err).Errorf("add tunnel error")
			time.Sleep(3 * time.Second)
			continue
		} else {
			logrus.Infof("add tunnel success")
			gost.NewGostClient().PostConfig()
			return nil
		}
	}
	return err
}

func Delete(tunnelName string, workers []*entities.Worker,
	allWorkers, allNodes map[string]string) error {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
		}
	}()

	oldNodes := forwardNodesTemplate(workers)
	newNodes := lo.Filter(oldNodes, func(node *config.ForwardNodeConfig, i int) bool {
		return node.Name != tunnelName
	})
	newServConf := serviceConfigTemplate(newNodes)

	maxRetry := 10
	var err error
	for i := 0; i < maxRetry; i++ {
		if err = gost.NewGostClient().PutService(VorkerDefaultServiceName, newServConf); err != nil {
			logrus.WithError(err).Errorf("delete tunnel error")
			time.Sleep(3 * time.Second)
			continue
		} else {
			logrus.Infof("delete tunnel success")
		}
	}
	if conf.AppConfigInstance.RunMode == "agent" {
		return err
	}

	rules := genIngressRules(allWorkers, allNodes)
	newRules := lo.Filter(rules, func(rule *config.IngressRuleConfig, i int) bool {
		return rule.Hostname != tunnelName
	})
	ingressConf := ingressConfigTemplate(newRules)
	for i := 0; i < maxRetry; i++ {
		if err = gost.NewGostClient().PutIngress(VorkerDefaultIngressName, ingressConf); err != nil {
			logrus.WithError(err).Errorf("add tunnel error")
			time.Sleep(3 * time.Second)
			continue
		} else {
			logrus.Infof("add tunnel success")
			gost.NewGostClient().PostConfig()
			return nil
		}
	}
	return err
}

func RelayServerRun() {
	p := &gostInstance{}
	if err := svc.Run(p); err != nil {
		log.Fatal(err)
	}
}
