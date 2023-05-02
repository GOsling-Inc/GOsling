package handlers

import (
	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/labstack/echo/v4"
)

type ManagerHandler struct {
	middleware *middleware.Middleware
}

func NewManagerHandler(m *middleware.Middleware) *ManagerHandler {
	return &ManagerHandler{
		middleware: m,
	}
}

func (h *ManagerHandler) Confirm(c echo.Context) error {
	manager_id := h.middleware.Auth(c.Request().Header)
	if manager_id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	operation_id := c.FormValue("Id")
	table := c.FormValue("Value")
	status := c.FormValue("Status")
	code, err := h.middleware.Confirm(operation_id, table, status)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, nil)
}
