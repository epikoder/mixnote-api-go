package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mixnote/mixnote-api-go/configs"
)

type home struct {}

func Home() (h *home) {
	return
}

func (*home) Index(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Welcome to " + configs.App.Name)
}