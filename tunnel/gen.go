package tunnel

import (
	"bytes"
	"html/template"
	"os"
	"vorker/conf"
	"vorker/utils"

	"github.com/sirupsen/logrus"
)

const (
	gostConfFilePath = "gost.yaml"
	gostConfTemplate = `{{ if eq .runMode "master" }}
services:
  - name: vorker-entry
    addr: :18080
    handler:
      type: relay
      auther: internal
      metadata:
        entryPoint: ":10080"
        ingress: vorker-ingress-default
    listener:
      type: ws
{{ end }}
api:
  addr: :7788
  pathPrefix: /
  accesslog: true
  auther: internal

authers:
  - name: internal
    auths:
      - username: {{ .username }}
        password: {{ .password }}`
)

func buildGostConf(username, password, runMode string) string {
	writer := new(bytes.Buffer)
	capTemplate := template.New("gostconf")
	capTemplate, err := capTemplate.Parse(gostConfTemplate)
	if err != nil {
		panic(err)
	}
	capTemplate.Execute(writer, map[string]string{
		"username": username,
		"password": password,
		"runMode":  runMode,
	})
	return writer.String()
}

func init() {
	if conf.AppConfigInstance.RunMode == "agent" {
		logrus.Infof("run in agent mode")
	} else {
		logrus.Infof("run in master mode")
	}
	_, err := os.Stat(gostConfFilePath)
	if err == nil {
		logrus.Infof("config file already exsit, remove")
		err = os.Remove(gostConfFilePath)
		if !os.IsNotExist(err) && err != nil {
			logrus.Panic(err)
		}
	}

	logrus.Infof("generate new one to /bin/gost.yml")
	// gen gost.yaml
	err = utils.WriteFile(gostConfFilePath,
		buildGostConf(conf.AppConfigInstance.TunnelUsername,
			conf.AppConfigInstance.TunnelPassword, conf.AppConfigInstance.RunMode))
	if err != nil {
		logrus.Panic(err)
	}
}
