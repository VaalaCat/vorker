package workerd

import (
	"os"
	"testing"
	"voker/entities"
)

func TestRun(t *testing.T) {
	c := BuildCapfile(
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
	Run("/workspaces/vorker/workerd/", []string{})
	ch := make(chan struct{})
	<-ch
}
