package handlers

import (
	"encoding/json"

	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	POST_User(echo.Context) error
	POST_Change_Main(echo.Context) error
	POST_Change_Password(echo.Context) error
}

type UserHandler struct {
	middleware middleware.IMiddleware
}

func NewUserHandler(m middleware.IMiddleware) *UserHandler {
	return &UserHandler{
		middleware: m,
	}
}

func (h *UserHandler) POST_User(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	code, user, err := h.middleware.GetUser(id)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{OBJ{
		"Name":      user.Name,
		"Surname":   user.Surname,
		"Email":     user.Email,
		"Birthdate": user.Birthdate,
	}, ""})
}

func (h *UserHandler) POST_Change_Main(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	temp_user := models.User{
		Id:        id,
		Name:      t["Name"].(string),
		Surname:   t["Surname"].(string),
		Birthdate: t["Birthdate"].(string),
	}

	code, err := h.middleware.Change_Main_Info(temp_user)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{"ok", ""})
}

func (h *UserHandler) POST_Change_Password(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	newPassword := t["NewPassword"].(string)
	oldPassword := t["OldPassword"].(string)

	code, err := h.middleware.Change_Password(id, oldPassword, newPassword)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{"ok", ""})
}
