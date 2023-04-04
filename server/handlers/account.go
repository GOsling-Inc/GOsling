package handlers

import (
	"errors"
	"strconv"
	"strings"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/labstack/echo/v4"
)

type IAccountHadler interface {
	POST_Add_Account(echo.Context) error
	POST_User_Accounts(echo.Context) error
	POST_Transfer(echo.Context) error
}

type AccountHandler struct {
	service *services.Service
}

func NewAccountHandler(s *services.Service) *AccountHandler {
	return &AccountHandler{
		service: s,
	}
}

func (h *AccountHandler) POST_User_Accounts(c echo.Context) error {
	header := c.Request().Header
	id, err := h.service.ParseJWT(header["Token"][0])
	if err != nil {
		return c.JSON(401, err.Error())
	}
	user, err := h.service.GetUser(id)
	if err != nil {
		return c.JSON(401, err.Error())
	}
	accs, err := h.service.GetUserAccounts(user)
	if err != nil {
		return c.JSON(401, err.Error())
	}
	return c.JSON(200, accs)
}

func (h *AccountHandler) POST_Add_Account(c echo.Context) error {
	beta_acc := &models.Account{
		Name: c.FormValue("Name"),
		Unit: c.FormValue("Unit"),
		Type: c.FormValue("Type"),
	}

	header := c.Request().Header
	id, err := h.service.ParseJWT(header["Token"][0])
	if err != nil {
		return c.JSON(401, err.Error())
	}
	user, err := h.service.GetUser(id)
	if err != nil {
		return c.JSON(401, err.Error())
	}
	if err = h.service.AddAccount(user, beta_acc); err != nil {
		return c.JSON(401, err.Error())
	}
	return nil
}

func (h *AccountHandler) POST_Transfer(c echo.Context) error {
	beta_trans := &models.Trasfer{
		Sender:   c.FormValue("Sender"),
		Receiver: c.FormValue("Receiver"),
	}
	beta_trans.Amount, _ = strconv.ParseFloat(c.FormValue("Amount"), 64)
	header := c.Request().Header
	id, err := h.service.ParseJWT(header["Token"][0])
	if err != nil {
		return c.JSON(401, err.Error())
	}
	if !strings.Contains(beta_trans.Sender, id) {
		return c.JSON(401, errors.New("uninitialized sender"))
	}
	h.service.ProvideTransfer(beta_trans)
	return nil
}
