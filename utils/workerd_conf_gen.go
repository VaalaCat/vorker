package utils

import (
	"bytes"
	"errors"
	"html/template"
	"path/filepath"
	"vorker/conf"
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

func GenWorkerConfig(worker *entities.Worker) error {
	if worker == nil || worker.GetUID() == "" {
		return errors.New("error worker")
	}
	fileMap := BuildCapfile([]*entities.Worker{
		worker,
	})

	fileContent, ok := fileMap[worker.GetUID()]
	if !ok {
		return errors.New("BuildCapfile error")
	}

	return WriteFile(
		filepath.Join(
			conf.AppConfigInstance.WorkerdDir,
			defs.WorkerInfoPath,
			worker.GetUID(),
			defs.CapFileName,
		), fileContent)
}
