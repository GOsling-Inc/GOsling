package handlers

import (
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service *services.Service
}

func (h *UserHandler) POST_User(c echo.Context) error {
	user := models.User{
		Name:      c.FormValue("Name"),
		Surname:   c.FormValue("Surname"),
		Birthdate: c.FormValue("Birthdate"),
	}
	header := c.Request().Header
	id, err := h.service.ParseJWT(header["Token"][0])
	if err != nil {
		c.JSON(401, err.Error())
	}
	if err = h.service.GetUser(id); err != nil {
		return c.JSON(401, err.Error())
	}
	return c.JSON(201, map[string]string{
		"Name":     user.Name,
		"Surname":  user.Surname,
		"Bithdate": user.Birthdate,
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
		c.JSON(401, err.Error())
	}
	if err = h.service.GetUser(id); err != nil {
		return c.JSON(401, err.Error())
	}
	temp_user.Id = id
	if err = h.service.Change_Main_Info(temp_user); err != nil {
		return c.JSON(500, err.Error())
	}
	return nil
}

func (h *UserHandler) POST_Change_Password(c echo.Context) error {
	temp_user := models.User{
		Password: c.FormValue("Password"),
	}
	header := c.Request().Header
	id, err := h.service.ParseJWT(header["Token"][0])
	if err != nil {
		c.JSON(401, err.Error())
	}
	if err = h.service.GetUser(id); err != nil {
		return c.JSON(401, err.Error())
	}
	temp_user.Id = id
	if err = h.service.Change_Password(temp_user); err != nil {
		return c.JSON(500, err.Error())
	}
	return nil
}
