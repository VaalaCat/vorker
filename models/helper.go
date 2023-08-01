package models

import "github.com/sirupsen/logrus"

func GetIngressParam() (
	allTunnel map[string]string,
	allNodes map[string]string,
) {
	var err error
	allTunnel, err = AdminGetAllWorkersTunnelMap()
	if err != nil {
		logrus.Errorf("get all workers failed: %v", err)
		return
	}

	allNodes, err = AdminGetAllNodesMap()
	if err != nil {
		logrus.Errorf("get all nodes map failed: %v", err)
		return
	}
	return
}
