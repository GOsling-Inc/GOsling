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
		Name:    c.FormValue("Name"),
		Surname: c.FormValue("Surname"),
		//Password: c.FormValue("Password"), //надо разделить на отдельную фунцию
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
		"Name":    user.Name,
		"Surname": user.Surname,
		//"Password": user.Password,
	})
}

func (h *UserHandler) POST_Change(c echo.Context) error {
	temp_user := models.User{
		Name:     c.FormValue("Name"),
		Surname:  c.FormValue("Surname"),
		Email:    c.FormValue("Email"),
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
	//update user info from temp_user by the id
	return nil
}
