package workerd

import (
	"vorker/conf"
	"vorker/defs"
	"vorker/entities"
	"vorker/models"
	"vorker/rpc"
	"vorker/utils"

	"github.com/google/uuid"
	"github.com/lucasepe/codename"
	"github.com/sirupsen/logrus"
)

func FillWorkerValue(worker *entities.Worker, keepUID bool, UID string, UserID uint) {
	if !keepUID {
		worker.UID = utils.GenerateUID()
	}
	if len(worker.TunnelID) == 0 || !keepUID {
		worker.TunnelID = uuid.New().String()
	}
	worker.UserID = uint64(UserID)
	worker.HostName = defs.DefaultHostName

	if len(worker.NodeName) == 0 {
		assignNode, err := models.GetAssignNode()
		if err == nil {
			worker.NodeName = assignNode.GetName()
		} else {
			worker.NodeName = defs.DefaultNodeName
		}
	}
	worker.ExternalPath = defs.DefaultExternalPath
	port, err := utils.GetAvailablePort(defs.DefaultHostName)
	if err != nil {
		logrus.Panic("get available port failed", err)
	}
	worker.Port = int32(port)

	if len(worker.Code) == 0 {
		worker.Code = []byte(defs.DefaultCode)
	}
	if len(worker.Entry) == 0 {
		worker.Entry = defs.DefaultEntry
	}
	// if the worker name is not unique, use the uid as the name
	if wl, err :=
		models.AdminGetWorkersByNames([]string{worker.Name}); len(wl) > 0 ||
		err != nil ||
		len(worker.Name) == 0 {
		if len(wl) == 1 {
			if UID == wl[0].UID {
				return
			}
		}
		rng, _ := codename.DefaultRNG()
		worker.Name = codename.Generate(rng, 0)
	}
}

func SyncAgent(w *entities.Worker) {
	go func(worker *entities.Worker) {
		if worker.NodeName == conf.AppConfigInstance.NodeName {
			return
		}

		targetNode, err := models.GetNodeByNodeName(worker.NodeName)
		if err != nil {
			logrus.Errorf("worker node is invalid, db error: %v", err)
			return
		}
		if err := rpc.EventNotify(targetNode, defs.EventSyncWorkers); err != nil {
			logrus.Errorf("emit event: %v error, err: %v", defs.EventSyncWorkers, err)
			return
		}
	}(w)
}
