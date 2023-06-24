package workerd

import (
	"path/filepath"
	"vorker/conf"
	"vorker/defs"
	"vorker/entities"
	"vorker/models"
	"vorker/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetConfByNodeNameEndpoint(c *gin.Context) {
	nodeName := c.GetString(defs.KeyNodeName)
	node, err := models.GetNodeByNodeName(nodeName)
	if err != nil || node == nil {
		logrus.Errorf("failed to get node by node name, err: %v", err)
		c.JSON(defs.CodeInternalError, gin.H{"message": err.Error()})
		return
	}

	workerList, err := models.AdminGetWorkersByNodeName(nodeName)
	if err != nil || len(workerList) == 0 {
		logrus.Errorf("failed to get all workers, err: %v", err)
		c.JSON(defs.CodeInternalError, gin.H{"message": err.Error()})
		return
	}

	capnp := utils.BuildCapfile(&entities.WorkerList{
		Workers: models.Trans2Entities(workerList),
	})

	if len(capnp) == 0 {
		logrus.Errorf("failed to build capnp file")
		c.JSON(defs.CodeInternalError, gin.H{"message": "failed to build capnp file"})
		return
	}

	c.JSON(defs.CodeSuccess, gin.H{"message": "success", "capnp": capnp})
}

func GenCapnpConfig() error {
	workerRecords, err := models.AdminGetWorkersByNodeName(conf.AppConfigInstance.NodeName)
	if err != nil {
		logrus.Errorf("failed to get all workers, err: %v", err)
	}

	workerList := &entities.WorkerList{
		Workers: models.Trans2Entities(workerRecords),
	}

	return utils.WriteFile(
		filepath.Join(
			conf.AppConfigInstance.WorkerdDir,
			defs.CapFileName,
		),
		utils.BuildCapfile(workerList))
}
