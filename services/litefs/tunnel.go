package litefs

import (
	"vorker/common"
	"vorker/conf"
	"vorker/tunnel"

	"github.com/sirupsen/logrus"
)

func InitTunnel() {
	if conf.IsMaster() {
		err := tunnel.GetClient().AddService(common.ServiceLitefs, conf.AppConfigInstance.LitefsPrimaryPort)
		if err != nil {
			logrus.WithError(err).Errorf("init tunnel for master litefs service error")
			return
		}
		logrus.Infof("init tunnel for litefs serivce success")
		return
	} else {
		err := tunnel.GetClient().AddVisitor(common.ServiceLitefs, conf.AppConfigInstance.LitefsPrimaryPort)
		if err != nil {
			logrus.WithError(err).Errorf("init tunnel for agent litefs visitor failed")
			return
		}
		logrus.WithError(err).Errorf("init tunnel for agent litefs visitor success")
		return
	}
}
