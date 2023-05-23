package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/labstack/echo/v4"
)

type IDepositHandler interface {
	POST_NewDeposit(echo.Context) error
	GET_User_Deposits(echo.Context) error
}

type DepositHandler struct {
	middleware middleware.IMiddleware
}

func NewDepositHandler(m middleware.IMiddleware) *DepositHandler {
	return &DepositHandler{
		middleware: m,
	}
}

func (h *DepositHandler) POST_NewDeposit(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	beta_depos := models.Deposit{
		UserId:    id,
		AccountId: t["AccountId"].(string),
		Period:    t["Period"].(string),
	}
	beta_depos.Amount, _ = strconv.ParseFloat(t["Amount"].(string), 64)
	beta_depos.Percent, _ = strconv.ParseFloat(t["Percent"].(string), 64)

	code, err := h.middleware.CreateDeposit(beta_depos)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{"ok", ""})
}

func (h *DepositHandler) GET_User_Deposits(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	code, depos, err := h.middleware.GetUserDeposits(id)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{depos, ""})
}
