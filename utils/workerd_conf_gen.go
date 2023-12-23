package utils

import (
	"bytes"
	"html/template"
	"vorker/defs"
	"vorker/entities"
)

func BuildCapfile(workers []*entities.Worker) map[string]string {
	if len(workers) == 0 {
		return map[string]string{}
	}

	results := map[string]string{}
	for _, worker := range workers {
		writer := new(bytes.Buffer)
		capTemplate := template.New("capfile")
		workerTemplate := worker.GetTemplate()
		if workerTemplate == "" {
			workerTemplate = defs.DefaultTemplate
		}

		capTemplate, err := capTemplate.Parse(workerTemplate)
		if err != nil {
			panic(err)
		}
		capTemplate.Execute(writer, worker)
		results[worker.GetUID()] = writer.String()
	}
	return results
}
