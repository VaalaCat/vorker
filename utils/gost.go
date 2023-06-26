package utils

import (
	"bytes"
	"html/template"
	"vorker/conf"
)

var (
	gostConfTemplate = `services:
  - name: vorker-entry
    addr: :18080
    handler:
      type: relay
      auther: internal
      metadata:
        entryPoint: ":10080"
        ingress: vorker-ingress
    listener:
      type: ws

ingresses:
  - name: vorker-ingress
    reload: 10s
    http:
      url: http://127.0.0.1:8888/api/agent/ingress

authers:
  - name: internal
    auths:
      - username: {{ .username }}
        password: {{ .password }}`
)

func buildGostConf(username, password string) string {
	writer := new(bytes.Buffer)
	capTemplate := template.New("gostconf")
	capTemplate, err := capTemplate.Parse(gostConfTemplate)
	if err != nil {
		panic(err)
	}
	capTemplate.Execute(writer, map[string]string{
		"username": username,
		"password": password,
	})
	return writer.String()
}

func init() {
	// gen gost.yaml
	WriteFile("gost.yaml",
		buildGostConf(conf.AppConfigInstance.TunnelUsername,
			conf.AppConfigInstance.TunnelPassword))
}
