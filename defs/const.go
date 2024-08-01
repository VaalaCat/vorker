package defs

const (
	CapFileName    = "workerd.capnp"
	WorkerInfoPath = "workers"
	WorkerCodePath = "src"
	DBTypeSqlite   = "sqlite"

	DefaultHostName     = "localhost"
	DefaultNodeName     = "default"
	DefaultExternalPath = "/"
	DefaultEntry        = "entry.js"
	DefaultCode         = `export default {
  async fetch(req, env) {
    try {
		let resp = new Response("worker: " + req.url + " is online! -- " + new Date())
		return resp
	} catch(e) {
		return new Response(e.stack, { status: 500 })
	}
  }
};`

	DefaultTemplate = `using Workerd = import "/workerd/workerd.capnp";

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
  modules = [
    (name = "{{.Entry}}", esModule = embed "src/{{.Entry}}"),
  ],
  compatibilityDate = "2023-04-03",
);`
)

const (
	KeyNodeName    = "node_name"
	KeyNodeSecret  = "node_secret"
	KeyNodeProto   = "node_proto"
	KeyWorkerProto = "worker_proto"
)

const (
	HeaderNodeName   = "x-node-name"
	HeaderNodeSecret = "x-secret"
	HeaderHost       = "Host"
)

const (
	CodeInvalidRequest = 400
	CodeUnAuthorized   = 401
	CodeNotFound       = 404
	CodeInternalError  = 500
	CodeSuccess        = 200
)

const (
	EventSyncWorkers  = "sync-workers"
	EventAddWorker    = "add-worker"
	EventDeleteWorker = "delete-worker"
	EventFlushWorker  = "flush-worker"
)
