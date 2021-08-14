package api

import (
	"github.com/mixnote/mixnote-api-go/src/core/service/totem/totem_middlewares"
	api_handlers "github.com/mixnote/mixnote-api-go/src/framework/handlers/api"
)

func registerJWTRoutes() {
	jwtGroup := echo_.Group("access")
	jwtGroup.GET("/user", api_handlers.JWT().User, totem_middlewares.JWTGuard)
	jwtGroup.GET("/refresh", api_handlers.JWT().Refresh, totem_middlewares.JWTRefresh)
}
