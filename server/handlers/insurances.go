package handlers

import (
	"strconv"

	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/labstack/echo/v4"
)

type EnsuranceHandler struct {
	middleware *middleware.Middleware
}

func NewEnsuranceHandler(m *middleware.Middleware) *EnsuranceHandler {
	return &EnsuranceHandler{
		middleware: m,
	}
}

func (h *EnsuranceHandler) POST_NewEnsurance(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	beta_ensure := models.Insurance{
		UserId:    id,
		AccountId: c.FormValue("AccountId"),
		Period:    c.FormValue("Period"),
	}
	beta_ensure.Amount, _ = strconv.ParseFloat(c.FormValue("Amount"), 64)
	code, err := h.middleware.CreateInsurance(beta_ensure)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{"ok", ""})
}

func (h *DepositHandler) GET_User_Ensurances(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	code, ensure, err := h.middleware.GetUserInsurances(id)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{ensure, ""})
}
