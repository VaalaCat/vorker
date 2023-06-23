package main

import (
	"vorker/entities"
	"vorker/models"

	"github.com/sirupsen/logrus"
)

func InitCache() {
	proxy := entities.GetProxy()
	tunnel := entities.GetTunnel()
	workerRecords, err := models.AdminGetAllWorkers()
	if err != nil {
		logrus.Errorf("failed to get all workers, err: %v", err)
	}
	workerList := &entities.WorkerList{
		Workers: models.Trans2Entities(workerRecords),
	}

	if err != nil {
		logrus.Errorf("failed to get all nodes, err: %v", err)
	}

	proxy.InitProxyMap(workerList)
	tunnel.InitTunnelMap(workerList)
}
