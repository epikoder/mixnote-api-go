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
	"github.com/mixnote/mixnote-api-go/src/framework/middlewares"
	"github.com/mixnote/mixnote-api-go/src/framework/router/routes/api"
	"github.com/mixnote/mixnote-api-go/src/framework/router/routes/web"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

var (
	server_host string = "mixnote.com"
	server_port int
	err         error

	apiURL, webURL string
	Domains        map[string]*Host = make(map[string]*Host)
)

func init() {
	fn := func() string {
		host_ := os.Getenv("SERVER_HOST")
		if host_ != "" {
			return host_
		}
		return server_host
	}
	server_host = fn()
	server_port, err = strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		server_port = 8080
	}
}

func Serve(_host string, _port int) {
	if _host != "" {
		server_host = _host
	}

	if strconv.Itoa(_port) != "" {
		server_port = _port
	}
	apiURL = "api." + server_host + ":" + fmt.Sprintf("%d", server_port)
	webURL = server_host + ":" + fmt.Sprintf("%d", server_port)
	Domains[apiURL] = &Host{echo.New()}
	Domains[webURL] = &Host{echo.New()}

	Domains[webURL].Echo.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("APP_KEY")))))
	Domains[webURL].Echo.Use(middleware.CORS())

	Domains[apiURL].Echo.Use(middlewares.JsonResponse)

	echo_ := echo.New()
	if configs.App.DEBUG {
		echo_.Use(middleware.Logger())
	}

	echo_.HideBanner = true
	echo_.IPExtractor = echo.ExtractIPFromXFFHeader(
		echo.TrustLinkLocal(false),
		// echo.TrustIPRange(&net.IPNet{}),
	)
	api.RegisterRoutes(Domains[apiURL].Echo)
	web.RegisterRoutes(Domains[webURL].Echo)

	domains := func(ctx echo.Context) (err error) {
		host := Domains[ctx.Request().Host]
		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(ctx.Response(), ctx.Request())
		}
		return
	}
	echo_.Any("/*", domains)
	echo_.Logger.Fatal(echo_.Start(fmt.Sprintf("%s:%d", server_host, server_port)))
}
