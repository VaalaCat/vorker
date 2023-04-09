package defs

const (
	CapFileName  = "workerd.capnp"
	DBTypeSqlite = "sqlite"

	DefaultCode = `addEventListener("fetch", event => {
	event.respondWith(new Response("Hello World"));
});`
)
