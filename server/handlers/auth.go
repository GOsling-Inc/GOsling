package handlers

import (
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
	user := models.User{
		Name:      c.FormValue("Name"),
		Surname:   c.FormValue("Surname"),
		Email:     c.FormValue("Email"),
		Password:  c.FormValue("Password"),
		Role:      "user",
		Birthdate: c.FormValue("Birthdate"),
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
	user := models.User{
		Email:    c.FormValue("Email"),
		Password: c.FormValue("Password"),
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
