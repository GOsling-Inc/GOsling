package handlers

import (
	"strconv"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/GOsling-Inc/GOsling/utils"
	"github.com/labstack/echo/v4"
)

type IAccountHadler interface {
	POST_Add_Account(echo.Context) error
	GET_User_Accounts(echo.Context) error
	POST_Freeze_Account(echo.Context) error
	POST_Delete_Account(echo.Context) error
	POST_Transfer(echo.Context) error
	POST_User_Exchange(echo.Context) error
}

type AccountHandler struct {
	service *services.Service
}

func NewAccountHandler(s *services.Service) *AccountHandler {
	return &AccountHandler{
		service: s,
	}
}

func (h *AccountHandler) GET_User_Accounts(c echo.Context) error {
	header := c.Request().Header
	if header["Token"] == nil {
		return c.JSON(401, "Unauthorized")
	}
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

func (h *AccountHandler) POST_Freeze_Account(c echo.Context) error {
	accountId := c.FormValue("AccountId")
	password := c.FormValue("Password")

	header := c.Request().Header
	id, err := h.service.ParseJWT(header["Token"][0])
	if err != nil {
		return c.JSON(401, err.Error())
	}
	if accountId[7:14] != id {
		return c.JSON(401, "incorrect account")
	}
	user, err := h.service.GetUser(id)
	if err != nil {
		return c.JSON(401, err.Error())
	}
	password, err = h.service.Hash(password)
	if err != nil {
		return c.JSON(401, err.Error())
	}
	if user.Password != password {
		return c.JSON(401, "incorrect password")
	}
	if err = h.service.FreezeAccount(accountId); err != nil {
		return c.JSON(401, err.Error())
	}
	return nil
}

func (h *AccountHandler) POST_Delete_Account(c echo.Context) error {
	accountId := c.FormValue("AccountId")
	password := c.FormValue("Password")

	header := c.Request().Header
	id, err := h.service.ParseJWT(header["Token"][0])
	if err != nil {
		return c.JSON(401, err.Error())
	}
	if accountId[7:14] != id {
		return c.JSON(401, "incorrect account")
	}
	user, err := h.service.GetUser(id)
	if err != nil {
		return c.JSON(401, err.Error())
	}
	password, err = h.service.Hash(password)
	if err != nil {
		return c.JSON(401, err.Error())
	}
	if user.Password != password {
		return c.JSON(401, "incorrect password")
	}
	if err = h.service.DeleteAccount(accountId); err != nil {
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
	if beta_trans.Sender[7:14] != id {
		return c.JSON(401, "incorrect account")
	}
	if err = h.service.ProvideTransfer(beta_trans); err != nil {
		return c.JSON(401, err.Error())
	}
	return err
}

func (h *AccountHandler) POST_User_Exchange(c echo.Context) error {
	beta_exchng := &models.Exchange{
		Sender:   c.FormValue("Sender"),
		Receiver: c.FormValue("Receiver"),
	}
	beta_exchng.SenderAmount, _ = strconv.ParseFloat(c.FormValue("Sender_amount"), 64)
	header := c.Request().Header
	id, err := h.service.ParseJWT(header["Token"][0])
	if err != nil {
		return c.JSON(401, err.Error())
	}
	if beta_exchng.Sender[7:14] != id || beta_exchng.Receiver[7:14] != id {
		return c.JSON(401, "incorrect account")
	}
	if err = h.service.ProvideExchange(beta_exchng); err != nil {
		return c.JSON(401, err.Error())
	}
	return nil
}

func (h *Handler) GET_Exchanges(c echo.Context) error {
	return c.JSON(200, map[string]float64{
		"BYN/USD": utils.BYN_USD(),
		"BYN/EUR": utils.BYN_EUR(),
	})
}
