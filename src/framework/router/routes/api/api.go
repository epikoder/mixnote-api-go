package api

import (
	"github.com/labstack/echo/v4"
	"github.com/mixnote/mixnote-api-go/src/core/service/totem/totem_middlewares"
	api_handlers "github.com/mixnote/mixnote-api-go/src/framework/handlers/api"
	"github.com/mixnote/mixnote-api-go/src/music/handlers/api"
)

var (
	echo_ *echo.Echo
)

func RegisterRoutes(e *echo.Echo) {
	echo_ = e
	registerJWTRoutes()
	
	registerAuthRoutes()
	registerAccountRoutes()
	registerMusicRoutes()
}

func registerJWTRoutes() {
	jwtGroup := echo_.Group("access")
	jwtGroup.GET("/user", api_handlers.JWT().User, totem_middlewares.JWTGuard)
	jwtGroup.GET("/refresh", api_handlers.JWT().Refresh, totem_middlewares.JWTRefresh)
}

func registerAuthRoutes() {
	authGroup := echo_.Group("auth")
	authGroup.POST("/login_email", api_handlers.Auth().LoginEmail).Name = "auth.login"
	authGroup.POST("/register", api_handlers.Auth().RegisterEmail).Name = "auth.register"
}

func registerAccountRoutes() {
	acc := echo_.Group("account")
	acc.GET("/enable2fa", api_handlers.Enabled2FA, totem_middlewares.JWTGuard)
}

func registerMusicRoutes() {
	muzz := echo_.Group("music")
	muzz.POST("/upload", music_api_handler.Upload, totem_middlewares.JWTGuard)
}

