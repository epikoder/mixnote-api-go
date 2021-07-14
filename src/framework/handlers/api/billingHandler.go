package api_handlers


import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	billing struct {}
)

func Billing() (a *billing) {
	return
}

func (*billing) Customer(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello Login page")
}

func (*billing) RegisterEmail(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"user": 1,
	})
}