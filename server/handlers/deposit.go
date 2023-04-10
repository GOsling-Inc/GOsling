package handlers

import (
	"strconv"

	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/labstack/echo/v4"
)

type DepositHandler struct {
	middleware *middleware.Middleware
}

func NewDepositHandler(m *middleware.Middleware) *DepositHandler {
	return &DepositHandler{
		middleware: m,
	}
}

func (h *DepositHandler) POST_NewDeposit(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	beta_depos := models.Deposit{
		UserId:    id,
		AccountId: c.FormValue("AccountId"),
		Period:    c.FormValue("Period"),
	}
	beta_depos.Amount, _ = strconv.ParseFloat(c.FormValue("Amount"), 64)
	beta_depos.Percent, _ = strconv.ParseFloat(c.FormValue("Percent"), 64)

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
