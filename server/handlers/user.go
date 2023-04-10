package handlers

import (
	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	middleware *middleware.Middleware
}

func NewUserHandler(m *middleware.Middleware) *UserHandler {
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

	temp_user := models.User{
		Id:        id,
		Name:      c.FormValue("Name"),
		Surname:   c.FormValue("Surname"),
		Birthdate: c.FormValue("Birthdate"),
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
	newPassword := c.FormValue("NewPassword")
	oldPassword := c.FormValue("OldPassword")

	code, err := h.middleware.Change_Password(id, oldPassword, newPassword)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{"ok", ""})
}
