package utils

import (
	"bytes"
	"html/template"
	"vorker/entities"
)

var capfileTemplate = `using Workerd = import "/workerd/workerd.capnp";

const config :Workerd.Config = (
  services = [
    (name = "{{.UID}}", worker = .v{{.UID}}Worker),
  ],

  sockets = [
    (
      name = "{{.UID}}",
      address = "{{.HostName}}:{{.Port}}",
      http=(),
      service="{{.UID}}"
    ),
  ]
);

const v{{.UID}}Worker :Workerd.Worker = (
  serviceWorkerScript = embed "src/{{.Entry}}",
  compatibilityDate = "2023-04-03",
);
`

func BuildCapfile(workers []*entities.Worker) map[string]string {
	if len(workers) == 0 {
		return map[string]string{}
	}

	results := map[string]string{}
	for _, worker := range workers {
		writer := new(bytes.Buffer)
		capTemplate := template.New("capfile")
		capTemplate, err := capTemplate.Parse(capfileTemplate)
		if err != nil {
			panic(err)
		}
		capTemplate.Execute(writer, worker)
		results[worker.GetUID()] = writer.String()
	}
	return results
}
