package tunnel

import (
	"vorker/conf"

	"github.com/VaalaCat/tunnel/server"
)

func Serve() {
	server.RunServer(conf.AppConfigInstance.RPCPort)
}
