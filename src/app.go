package app

import (
	"github.com/mixnote/mixnote-api-go/src/framework/router"
)

// Bootstrap
func init() {}

func StartServer(host string, port int) {
	server.Serve(host, port)
}
