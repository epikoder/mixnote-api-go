package web

import (
	"github.com/labstack/echo/v4"
	"github.com/mixnote/mixnote-api-go/src/framework/handlers"
	"github.com/mixnote/mixnote-api-go/src/framework/middlewares"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/home", handlers.Home().Index)
	e.Static("/static", "public/assets")
	e.File("/", "public/index.html", middlewares.NoCacheControl)
}
