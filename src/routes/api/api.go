package api

import (
	"github.com/labstack/echo/v4"
	"github.com/mixnote/mixnote-api-go/src/framework/handlers/api/auth"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/login", auth.LoginEmail)
}