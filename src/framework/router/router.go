package server

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mixnote/mixnote-api-go/configs"
	"github.com/mixnote/mixnote-api-go/src/framework/router/routes/api"
	"github.com/mixnote/mixnote-api-go/src/framework/router/routes/web"
)

var (
	host  string = "127.0.0.1"
	port  int
	err   error
	echo_ *echo.Echo = echo.New()
)

func init() {
	fn := func() string {
		host_ := os.Getenv("SERVER_HOST")
		if host_ != "" {
			return host_
		}
		return host
	}
	host = fn()
	port, err = strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		port = 8080
	}
}

func Serve(_host string, _port int) {
	if _host != "" {
		host = _host
	}

	if strconv.Itoa(_port) != "" {
		port = _port
	}

	hf := func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Add("content-type", "application/json")
			return hf(c)
		}
	}
	
	echo_.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("APP_KEY")))))
	echo_.Use(hf)
	echo_.Use(middleware.CORS())
	if configs.App.DEBUG {
		echo_.Use(middleware.Logger())
	}

	api.RegisterRoutes(echo_)
	web.RegisterRoutes(echo_)
	echo_.Logger.Fatal(echo_.Start(fmt.Sprintf("%s:%d", host, port)))
}
