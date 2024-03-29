package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/labstack/echo/v4"
)

type IAccountHandler interface {
	GET_User_Accounts(c echo.Context) error
	POST_Add_Account(c echo.Context) error
	GET_User_Transfers(c echo.Context) error
	POST_Delete_Account(c echo.Context) error
	POST_Transfer(c echo.Context) error
	POST_User_Exchange(c echo.Context) error
	GET_Exchanges(c echo.Context) error
}

type AccountHandler struct {
	middleware middleware.IMiddleware
}

func NewAccountHandler(m middleware.IMiddleware) *AccountHandler {
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

func (h *AccountHandler) GET_User_Transfers(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	code, trs := h.middleware.UserTransfers(id)
	return c.JSON(code, JSON{trs, ""})
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
	transfer.Amount, _ = strconv.ParseFloat(t["Amount"].(string), 64)

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
	exc.SenderAmount, _ = strconv.ParseFloat(t["Sender_Amount"].(string), 64)

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
