package app

import (
	"github.com/mixnote/mixnote-api-go/src/core/service/totem"
	"github.com/mixnote/mixnote-api-go/src/framework/router"
)

// Bootstrap
func init() {
	totem.UseRedis()
	totem.AccessExpiresIn(60*24*7) //TODO remove this
}

func StartServer(host string, port int) {
	server.Serve(host, port)
}
