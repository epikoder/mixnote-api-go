package api

import (
	// "github.com/mixnote/mixnote-api-go/src/core/service/totem/totem_middlewares"
	api_handlers "github.com/mixnote/mixnote-api-go/src/framework/handlers/api"
	// "github.com/mixnote/mixnote-api-go/src/framework/middlewares"
)

func registerAuthRoutes() {
	authGroup := echo_.Group("auth")
	authGroup.POST("/login_email", api_handlers.Auth().LoginEmail).Name = "auth.login"
	authGroup.POST("/register", api_handlers.Auth().RegisterEmail).Name = "auth.register"
}
