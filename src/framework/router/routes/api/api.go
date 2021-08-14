package api

import (
	"github.com/labstack/echo/v4"
)

var (
	echo_ *echo.Echo
)

func RegisterRoutes(e *echo.Echo) {
	echo_ = e
	registerAuthRoutes()
	registerJWTRoutes()
}
