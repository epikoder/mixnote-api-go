package api_handlers

import (
	"encoding/json"
	"net/http"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mixnote/mixnote-api-go/src/core/models"
	"github.com/mixnote/mixnote-api-go/src/framework/hasher"
	"github.com/mixnote/mixnote-api-go/src/framework/repositories"
	"github.com/mixnote/mixnote-api-go/src/framework/validator"
)

type (
	auth struct{}
)

var (
	u *models.User = &models.User{}
)

func Auth() (a *auth) {
	return
}

// API Route login
func (*auth) LoginEmail(ctx echo.Context) error {
	rules := map[string][]string{
		"email":    {"email"},
		"password": {"required", "alpha_numeric", "min:8"},
	}
	form, bag, err := validator.NewValidator(ctx.Request(), rules, nil).Validate()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	
	if bag.Failed() {
		return ctx.JSON(http.StatusBadRequest, bag)
	}

	if ok := repositories.UserRepository(u).FindByEmail(form.GetString("email")); !ok {
		return ctx.String(http.StatusUnauthorized, "user does not exist")
	}

	if ok := hasher.Check(u.UserCredential.Password, form.GetString("password")); !ok {
		return ctx.String(http.StatusUnauthorized, "password does not match")
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"user": u,
	})
}

func (*auth) RegisterEmail(ctx echo.Context) (err error) {
	rules := map[string][]string{
		"email":    {"email"},
		"password": {"required", "alpha_numeric", "min:8"},
	}
	form, bag, err := validator.NewValidator(ctx.Request(), rules, nil).Validate()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	
	if bag.Failed() {
		return ctx.JSON(http.StatusBadRequest, bag)
	}

	//
	if ok := repositories.UserRepository(u).FindByEmail(form.GetString("email")); ok {
		return ctx.String(http.StatusOK, "user already exist")
	}
	p, err := hasher.Make(form.GetString("password"))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Server error")
	}
	repositories.UserRepository(u).CreateUser(&models.User{
		ID:    uuid.New(),
		Email: form.GetString("email"),
		UserCredential: models.UserCredential{
			ID:       uuid.New(),
			Password: p,
		},
	})
	return ctx.String(http.StatusOK, (func() string {
		n, _ := json.Marshal(u)
		return string(n)
	})())
}
