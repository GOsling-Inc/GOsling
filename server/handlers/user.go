package handlers

import (
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service *services.Service
}

func NewUserHandler(s *services.Service) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (h *UserHandler) POST_User(c echo.Context) error {
	header := c.Request().Header
	id, err := h.service.ParseJWT(header["Token"][0])
	if err != nil {
		return c.JSON(401, err.Error())
	}
	user, err := h.service.GetUser(id)
	if err != nil {
		return c.JSON(401, err.Error())
	}
	return c.JSON(200, map[string]string{
		"Name":      user.Name,
		"Surname":   user.Surname,
		"Email":     user.Email,
		"Birthdate": user.Birthdate,
	})
}

func (h *UserHandler) POST_Change_Main(c echo.Context) error {
	temp_user := models.User{
		Name:      c.FormValue("Name"),
		Surname:   c.FormValue("Surname"),
		Birthdate: c.FormValue("Birthdate"),
	}
	header := c.Request().Header
	id, err := h.service.ParseJWT(header["Token"][0])
	if err != nil {
		return c.JSON(401, err.Error())
	}
	_, err = h.service.GetUser(id)
	if err != nil {
		return c.JSON(401, err.Error())
	}
	temp_user.Id = id
	if err = h.service.Change_Main_Info(temp_user); err != nil {
		return c.JSON(500, err.Error())
	}
	return nil
}

func (h *UserHandler) POST_Change_Password(c echo.Context) error {
	tempUser := models.User{
		Password: c.FormValue("NewPassword"),
	}
	oldPassword := c.FormValue("OldPassword")

	tempUser.Password, _ = h.service.Hash(tempUser.Password)
	oldPassword, _ = h.service.Hash(oldPassword)

	header := c.Request().Header
	id, err := h.service.ParseJWT(header["Token"][0])
	if err != nil {
		return c.JSON(401, err.Error())
	}
	user, err := h.service.GetUser(id)
	if err != nil {
		return c.JSON(401, err.Error())
	}
	if user.Password != oldPassword {
		return c.JSON(401, "wrong password")
	}
	tempUser.Id = id
	if err = h.service.Change_Password(tempUser); err != nil {
		return c.JSON(500, err.Error())
	}
	return nil
}
