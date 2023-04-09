package workerd

import (
	"fmt"
	"testing"
	"voker/entities"
)

func TestBuildCapfile(t *testing.T) {
	type args struct {
		wokers *entities.WorkerList
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "common case",
			args: args{
				wokers: &entities.WorkerList{
					Workers: []*entities.Worker{
						{
							UID:          "test",
							Port:         8080,
							HostName:     "localhost",
							NodeName:     "test",
							ExternalPath: "/teste",
							Entry:        "test/entry.js",
						},
						{
							UID:          "test1",
							Port:         8081,
							HostName:     "localhost",
							NodeName:     "test1",
							ExternalPath: "/teste1",
							Entry:        "test1/entry.js",
						},
					},
				},
			},
			want: `using Workerd = import "/workerd/workerd.capnp";

const config :Workerd.Config = (
  services = [
    (name = "test", worker = .testWorker),
    (name = "test1", worker = .test1Worker),
  ],

  sockets = [
    (
      name = "test",
      address = "localhost:8080",
      http=(),
      service="test"
    ),
    (
      name = "test1",
      address = "localhost:8081",
      http=(),
      service="test1"
    ),
  ]
);

const testWorker :Workerd.Worker = (
  serviceWorkerScript = embed "workers/test/entry.js",
  compatibilityDate = "2023-04-03",
);
const test1Worker :Workerd.Worker = (
  serviceWorkerScript = embed "workers/test1/entry.js",
  compatibilityDate = "2023-04-03",
);
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildCapfile(tt.args.wokers); got != tt.want {
				fmt.Print(got)
				t.Errorf("BuildCapfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
