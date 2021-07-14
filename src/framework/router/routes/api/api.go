package api

import (
	"github.com/labstack/echo/v4"
	api_handlers "github.com/mixnote/mixnote-api-go/src/framework/handlers/api"
)

func RegisterRoutes(e *echo.Echo) {
	authGroup := e.Group("auth")
	authGroup.GET("/login", api_handlers.Auth().LoginEmail).Name = "auth.login"
	authGroup.GET("/register", api_handlers.Auth().RegisterEmail).Name = "auth.register"

	authGroup = e.Group("billing")
	authGroup.POST("/customer", api_handlers.Billing().Customer).Name = "billing.customer"
}
