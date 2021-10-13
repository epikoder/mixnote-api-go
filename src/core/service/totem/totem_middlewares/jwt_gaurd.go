package totem_middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mixnote/mixnote-api-go/configs"
	"github.com/mixnote/mixnote-api-go/src/core/models"
	"github.com/mixnote/mixnote-api-go/src/core/service/totem/guard"
	"github.com/mixnote/mixnote-api-go/src/framework/cache"
	"github.com/mixnote/mixnote-api-go/src/framework/repositories"
	"github.com/mixnote/mixnote-api-go/src/framework/utilities"
)

func JWTGuard(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		bearerToken := extractToken(ctx)
		if bearerToken == "" {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Unautorized",
			})
		}

		token, err := verifyToken(bearerToken)
		if err != nil || token != nil && !token.Valid {
			message := "Unautorized"
			status := http.StatusUnauthorized
			if strings.Contains(err.Error(), "expired") {
				message = "expired"
				status = http.StatusBadRequest
			}
			return ctx.JSON(status, map[string]string{
				"message": message,
			})
		}

		claims, ok := extractMetaData(token)
		if !ok {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Unautorized",
			})
		}

		k, ok := (claims["aud"].(map[string]interface{}))["id"].(string)
		if !ok {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Unautorized",
			})
		}

		if !isUserExist(k) {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Unautorized",
			})
		}
		
		return next(ctx)
	}
}

func extractToken(ctx echo.Context) string {
	strArr := strings.Split(ctx.Request().Header.Get("Authorization"), " ")
	if len(strArr) != 2 {
		return ""
	}
	return strArr[1]
}

func verifyToken(s string) (*jwt.Token, error) {
	token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		JWT_SECRET := os.Getenv("JWT_SECRET")
		if JWT_SECRET == "" {
			utilities.Console.Fatal("JWT_SECRET not found")
		}

		privKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(JWT_SECRET))
		if err != nil {
			utilities.Console.Fatal("Could not generate private key")
		}
		return &privKey.PublicKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func extractMetaData(t *jwt.Token) (jwt.MapClaims, bool) {
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !validateMetaData(claims) {
		return nil, false
	}
	return claims, true
}

func validateMetaData(claims jwt.MapClaims) bool {
	cache.UseDB(1)
	id, ok := claims["id"].(string)
	if !ok {
		return false
	}

	userID, err := cache.Get(id)
	if err != nil {
		utilities.Console.Fatal(err)
	}
	aud, ok := claims["aud"].(map[string]interface{})
	if !ok {
		return false
	}

	return userID == aud["id"] && claims["iss"] == configs.App.Name
}

func isUserExist(userID string) bool {
	u := &models.User{}
	repositories.UserRepository(u).FindByID(userID)
	guard.SetUser(u)
	return u.ID != uuid.Nil && !u.DeletedAt.Valid
}
