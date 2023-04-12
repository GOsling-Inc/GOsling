package handlers

import (
	"strconv"

	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/labstack/echo/v4"
)

type InsuranceHandler struct {
	middleware *middleware.Middleware
}

func NewInsuranceHandler(m *middleware.Middleware) *InsuranceHandler {
	return &InsuranceHandler{
		middleware: m,
	}
}

func (h *InsuranceHandler) POST_NewInsurance(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	insurance := models.Insurance{
		UserId:    id,
		AccountId: c.FormValue("AccountId"),
		Period:    c.FormValue("Period"),
	}
	insurance.Amount, _ = strconv.ParseFloat(c.FormValue("Amount"), 64)
	code, err := h.middleware.CreateInsurance(insurance)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{"ok", ""})
}

func (h *InsuranceHandler) GET_User_Insurances(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	code, insurances, err := h.middleware.GetUserInsurances(id)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{insurances, ""})
}
