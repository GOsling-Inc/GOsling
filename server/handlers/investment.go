package handlers

import (
	"strconv"

	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/labstack/echo/v4"
)

type InvestmentHandler struct {
	middleware *middleware.Middleware
}

func NewInvestmentHandler(m *middleware.Middleware) *InvestmentHandler {
	return &InvestmentHandler{
		middleware: m,
	}
}

func (h *InvestmentHandler) POST_User_Stocks_NewOrder(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	new_order := models.Order{
		UserId:    id,
		AccountId: c.FormValue("AccountId"),
		Name:      c.FormValue("Name"),
		Action:    c.FormValue("Action"),
	}
	new_order.Price, _ = strconv.ParseFloat(c.FormValue("Price"), 64)
	new_order.Count, _ = strconv.Atoi(c.FormValue("Count"))
	code, err := h.middleware.CreateOrder(new_order)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{"ok", ""})
}

func (h *InvestmentHandler) POST_User_Stocks_Buy(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	code, err := h.middleware.BuyStock(c.FormValue("OrderId"), c.FormValue("AccountId"), c.FormValue("Count"), id)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{"ok", ""})
}

func (h *InvestmentHandler) POST_User_Stocks_Sell(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	code, err := h.middleware.SellStock(c.FormValue("OrderId"), c.FormValue("AccountId"), c.FormValue("Count"), id)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{"ok", ""})
}
