package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/labstack/echo/v4"
)

type IInvestmentHandler interface {
	GET_Orders(c echo.Context) error
	POST_User_Stocks_NewOrder(c echo.Context) error
	POST_User_Stocks_Buy(c echo.Context) error
	POST_User_Stocks_Sell(c echo.Context) error
}

type InvestmentHandler struct {
	middleware middleware.IMiddleware
}

func NewInvestmentHandler(m middleware.IMiddleware) *InvestmentHandler {
	return &InvestmentHandler{
		middleware: m,
	}
}

func (h *InvestmentHandler) GET_Orders(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	code, orders := h.middleware.Orders()
	return c.JSON(code, JSON{orders, ""})
}

func (h *InvestmentHandler) POST_User_Stocks_NewOrder(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	new_order := models.Order{
		UserId:    id,
		AccountId: t["AccountId"].(string),
		Name:      t["Name"].(string),
		Action:    t["Action"].(string),
	}
	new_order.Price, _ = strconv.ParseFloat(t["Price"].(string), 64)
	new_order.Count, _ = strconv.Atoi(t["Count"].(string))
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

	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	code, err := h.middleware.BuyStock(t["OrderId"].(string), t["AccountId"].(string), t["Count"].(string), id)
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

	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	code, err := h.middleware.SellStock(t["OrderId"].(string), t["AccountId"].(string), t["Count"].(string), id)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{"ok", ""})
}
