package api_handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	auth struct {}
)

func Auth() (a *auth) {
	return
}

func (*auth) LoginEmail(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello Login page")
}

func (*auth) RegisterEmail(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"user": 1,
	})
}