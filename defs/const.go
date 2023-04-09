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
