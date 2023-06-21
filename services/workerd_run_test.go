package services

import (
	"os"
	"testing"
	"vorker/entities"
	"vorker/utils"
)

func TestRun(t *testing.T) {
	c := utils.BuildCapfile(
		&entities.WorkerList{
			Workers: []*entities.Worker{
				{
					UID:          "test",
					Port:         8080,
					HostName:     "localhost",
					NodeName:     "test",
					ExternalPath: "/test",
					Entry:        "test/entry.js",
				},
			},
		},
	)
	f, err := os.Create("/workspaces/vorker/workerd/workerd.capnp")
	if err != nil {
		panic(err)
	}
	f.WriteString(c)
	WorkerdRun("/workspaces/vorker/workerd/", []string{})
	ch := make(chan struct{})
	<-ch
}
