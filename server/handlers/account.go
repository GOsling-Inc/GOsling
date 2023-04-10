package handlers

import (
	"strconv"

	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	middleware *middleware.Middleware
}

func NewAccountHandler(m *middleware.Middleware) *AccountHandler {
	return &AccountHandler{
		middleware: m,
	}
}

func (h *AccountHandler) GET_User_Accounts(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	code, accs, err := h.middleware.GetUserAccounts(id)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{accs, ""})
}

func (h *AccountHandler) POST_Add_Account(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	acc := models.Account{
		Name: c.FormValue("Name"),
		Unit: c.FormValue("Unit"),
		Type: c.FormValue("Type"),
	}

	code, err := h.middleware.AddAccount(id, acc)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{"ok", ""})
}

func (h *AccountHandler) POST_Delete_Account(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	accountId := c.FormValue("AccountId")
	password := c.FormValue("Password")

	code, err := h.middleware.DeleteAccount(id, accountId, password)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{"ok", ""})
}

func (h *AccountHandler) POST_Transfer(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	transfer := models.Trasfer{
		Sender:   c.FormValue("Sender"),
		Receiver: c.FormValue("Receiver"),
	}
	transfer.Amount, _ = strconv.ParseFloat(c.FormValue("Amount"), 64)

	code, err := h.middleware.ProvideTransfer(id, transfer)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{"ok", ""})
}

func (h *AccountHandler) POST_User_Exchange(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	exc := models.Exchange{
		Sender:   c.FormValue("Sender"),
		Receiver: c.FormValue("Receiver"),
	}
	exc.SenderAmount, _ = strconv.ParseFloat(c.FormValue("Sender_amount"), 64)

	code, err := h.middleware.ProvideExchange(id, exc)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{"ok", ""})
}

func (h *AccountHandler) GET_Exchanges(c echo.Context) error {
	return c.JSON(200, map[string]float64{
		"BYN/USD": h.middleware.BYN_USD(),
		"BYN/EUR": h.middleware.BYN_EUR(),
	})
}
