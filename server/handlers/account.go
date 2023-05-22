package handlers

import (
	"encoding/json"

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

	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	acc := models.Account{
		Name: t["Name"].(string),
		Unit: t["Unit"].(string),
		Type: t["Type"].(string),
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

	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	accountId := t["AccountId"].(string)
	password := t["Password"].(string)

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

	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	transfer := models.Trasfer{
		Sender:   t["Sender"].(string),
		Receiver: t["Receiver"].(string),
	}
	transfer.Amount, _ = t["Amount"].(float64)

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

	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	exc := models.Exchange{
		Sender:   t["Sender"].(string),
		Receiver: t["Receiver"].(string),
	}
	exc.SenderAmount, _ = t["Sender_amount"].(float64)

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
