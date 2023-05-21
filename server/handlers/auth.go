package handlers

import (
	"encoding/json"

	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/labstack/echo/v4"
)

type IAuthHandler interface {
	POST_SignUp(echo.Context) error
	POST_SignIn(echo.Context) error
}

type AuthHandler struct {
	middleware middleware.IMiddleware
}

func NewAuthHandler(m middleware.IMiddleware) *AuthHandler {
	return &AuthHandler{
		middleware: m,
	}
}

func (h *AuthHandler) POST_SignUp(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	user := models.User{
		Name:      t["Name"].(string),
		Surname:   t["Surname"].(string),
		Email:     t["Email"].(string),
		Password:  t["Password"].(string),
		Role:      "user",
		Birthdate: t["Birthdate"].(string),
	}

	code, err := h.middleware.SignUp(&user)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}

	token, jwtErr := h.middleware.CreateJWT(user.Id)
	if jwtErr != nil {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, err.Error()})
	}

	return c.JSON(code, JSON{OBJ{"Token": token}, ""})
}

func (h *AuthHandler) POST_SignIn(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	user := models.User{
		Email:    t["Email"].(string),
		Password: t["Password"].(string),
	}

	code, err := h.middleware.SignIn(&user)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}

	token, jwtErr := h.middleware.CreateJWT(user.Id)
	if jwtErr != nil {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, err.Error()})
	}

	return c.JSON(code, JSON{OBJ{"Token": token}, ""})
}
