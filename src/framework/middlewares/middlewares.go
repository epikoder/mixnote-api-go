package middlewares

import "github.com/labstack/echo/v4"

func JsonResponse(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Add("content-type", "application/json")
		return next(ctx)
	}
}

func NoCacheControl(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Add("Cache-Control", "no-cache")
		return next(ctx)
	}
}