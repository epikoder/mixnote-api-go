package server

import (
	"fmt"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mixnote/mixnote-api-go/src/routes/api"
)

var host string = "127.0.0.1"
var port int
var echo_ *echo.Echo = echo.New()

func init() {
	host = (func() string {
		host_ := os.Getenv("SERVER_HOST")
		if host_ != "" {
			return host_
		}
		return host
	})()

	var err error
	port, err = strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		port = 8080
	}
}

func GetRouter() (*echo.Echo) {
	return echo_
}

func Serve(_host string, _port int) {
	if _host != "" {
		host = _host
	}

	if strconv.Itoa(_port) != "" {
		port = _port
	}

	api.RegisterRoutes(echo_)
	echo_.Logger.Fatal(echo_.Start(fmt.Sprintf("%s:%d", host, port)))
}
