package web

import (
	"github.com/labstack/echo/v4"
	"github.com/mixnote/mixnote-api-go/src/framework/handlers"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("", handlers.Home().Index)
}