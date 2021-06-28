package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func init() {}

func LoginEmail(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello Login page")
}