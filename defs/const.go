package defs

const (
	CapFileName    = "workerd.capnp"
	WorkerCodePath = "workers"
	DBTypeSqlite   = "sqlite"

	DefaultHostName     = "localhost"
	DefaultNodeName     = "default"
	DefaultExternalPath = "/"
	DefaultEntry        = "entry.js"
	DefaultCode         = `addEventListener("fetch", event => {
	event.respondWith(new Response("Hello World"));
});`
)

const (
	KeyNodeName   = "node_name"
	KeyNodeSecret = "node_secret"
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
	EventSyncWorkers = "sync-workers"
)
