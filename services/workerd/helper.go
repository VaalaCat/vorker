package workerd

import (
	"voker/defs"
	"voker/entities"
	"voker/models"
	"voker/utils"

	"github.com/sirupsen/logrus"
)

func FillWorkerValue(worker *entities.Worker, keepUID bool, keepName bool) {
	if !keepUID {
		worker.UID = utils.GenerateUID()
	}
	worker.HostName = defs.DefaultHostName
	worker.NodeName = defs.DefaultNodeName
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
		models.GetWorkersByNames([]string{worker.Name}); (len(wl) > 0 ||
		err != nil ||
		len(worker.Name) == 0) && !keepName {
		worker.Name = worker.UID
	}
}
