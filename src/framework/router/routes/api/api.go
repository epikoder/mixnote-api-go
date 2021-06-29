package api

import (
	"github.com/labstack/echo/v4"
	api_handlers "github.com/mixnote/mixnote-api-go/src/framework/handlers/api"
)

func RegisterRoutes(e *echo.Echo) {
	authGroup := e.Group("auth")
	authGroup.GET("/login", api_handlers.Auth().LoginEmail)
}