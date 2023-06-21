package utils

import (
	"bytes"
	"html/template"
	"vorker/entities"
)

var capfileTemplate = `using Workerd = import "/workerd/workerd.capnp";

const config :Workerd.Config = (
  services = [
{{range $i, $worker := .Workers}}    (name = "{{$worker.UID}}", worker = .v{{$worker.UID}}Worker),
{{end}}  ],

  sockets = [{{range $i, $worker := .Workers}}
    (
      name = "{{$worker.UID}}",
      address = "{{$worker.HostName}}:{{$worker.Port}}",
      http=(),
      service="{{$worker.UID}}"
    ),{{end}}
  ]
);
{{range $i, $worker := .Workers}}
const v{{$worker.UID}}Worker :Workerd.Worker = (
  serviceWorkerScript = embed "workers/{{$worker.UID}}/{{$worker.Entry}}",
  compatibilityDate = "2023-04-03",
);{{end}}
`

func BuildCapfile(workers *entities.WorkerList) string {
	writer := new(bytes.Buffer)
	if len(workers.GetWorkers()) == 0 {
		return ""
	}
	capTemplate := template.New("capfile")
	capTemplate, err := capTemplate.Parse(capfileTemplate)
	if err != nil {
		panic(err)
	}
	capTemplate.Execute(writer, workers)
	return writer.String()
}
