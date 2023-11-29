package utils

import (
	"testing"
	"vorker/entities"

	"github.com/stretchr/testify/assert"
)

func TestBuildCapfile(t *testing.T) {
	tests := []struct {
		name   string
		wokers []*entities.Worker
		expect func(t *testing.T, resp map[string]string)
	}{
		{
			name: "common case",
			wokers: []*entities.Worker{
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

			expect: func(t *testing.T, result map[string]string) {
				assert.Equal(t,
					`using Workerd = import "/workerd/workerd.capnp";

const config :Workerd.Config = (
  services = [
    (name = "test", worker = .vtestWorker),
  ],

  sockets = [
    (
      name = "test",
      address = "localhost:8080",
      http=(),
      service="test"
    ),
  ]
);

const vtestWorker :Workerd.Worker = (
  serviceWorkerScript = embed "src/test/entry.js",
  compatibilityDate = "2023-04-03",
);
`, result["test"])

				assert.Equal(t, `using Workerd = import "/workerd/workerd.capnp";

const config :Workerd.Config = (
  services = [
    (name = "test1", worker = .vtest1Worker),
  ],

  sockets = [
    (
      name = "test1",
      address = "localhost:8081",
      http=(),
      service="test1"
    ),
  ]
);

const vtest1Worker :Workerd.Worker = (
  serviceWorkerScript = embed "src/test1/entry.js",
  compatibilityDate = "2023-04-03",
);
`, result["test1"])
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expect(t, BuildCapfile(tt.wokers))
		})
	}
}
