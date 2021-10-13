package totem_middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/mixnote/mixnote-api-go/src/framework/cache"
	"github.com/mixnote/mixnote-api-go/src/framework/utilities"
)

func JWTRefresh(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		message := "Unautorized"
		bearerToken := extractToken(ctx)
		if bearerToken == "" {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{
				"message": message,
			})
		}

		token, err := verifyToken(bearerToken)
		if err != nil || token != nil && !token.Valid {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{
				"message": message,
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Unautorized",
			})
		}

		id, ok := claims["id"].(string)
		accessId, ok_ := claims["aid"].(string)
		if !ok || !ok_ {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Unautorized",
			})
		}

		cache.UseDB(1)
		userId, err := cache.Pull(id)
		if err != nil {
			utilities.Console.Fatal(err)
		}

		cache.Pull(accessId)
		if !isUserExist(userId) {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Unautorized",
			})
		}

		return next(ctx)
	}
}