package workerd

import (
	"voker/defs"
	"voker/entities"
	"voker/models"
	"voker/utils"

	"github.com/google/uuid"
	"github.com/lucasepe/codename"
	"github.com/sirupsen/logrus"
)

func FillWorkerValue(worker *entities.Worker, keepUID bool, UID string) {
	if !keepUID {
		worker.UID = utils.GenerateUID()

	}
	if len(worker.TunnelID) == 0 || !keepUID {
		worker.TunnelID = uuid.New().String()
	}

	worker.HostName = defs.DefaultHostName
	if len(worker.NodeName) == 0 {
		worker.NodeName = defs.DefaultNodeName
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
