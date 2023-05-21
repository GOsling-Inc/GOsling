package handlers

import (
	"encoding/json"

	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/labstack/echo/v4"
)

type IManagerHandler interface {
	GetConfirms(echo.Context) error
	Confirm(echo.Context) error
	GetAccounts(echo.Context) error
	FreezeAccount(echo.Context) error
	BlockAccount(echo.Context) error
	GetTransactions(c echo.Context) error
	CancelTransaction(c echo.Context) error
	GetUsers(c echo.Context) error
	UpdateRole(c echo.Context) error
}

type ManagerHandler struct {
	middleware middleware.IMiddleware
}

func NewManagerHandler(m middleware.IMiddleware) *ManagerHandler {
	return &ManagerHandler{
		middleware: m,
	}
}

func (h *ManagerHandler) GetConfirms(c echo.Context) error {
	manager_id := h.middleware.Auth(c.Request().Header)
	if manager_id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	if err := h.middleware.AuthManager(manager_id); err != nil {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	code, confirms := h.middleware.GetConfirms()
	return c.JSON(code, JSON{confirms, ""})
}

func (h *ManagerHandler) Confirm(c echo.Context) error {
	manager_id := h.middleware.Auth(c.Request().Header)
	if manager_id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	if err := h.middleware.AuthManager(manager_id); err != nil {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	operation_id := t["Id"].(string)
	table := t["Value"].(string)
	status := t["Status"].(string)

	code, err := h.middleware.Confirm(operation_id, table, status)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, nil)
}

func (h *ManagerHandler) GetAccounts(c echo.Context) error {
	manager_id := h.middleware.Auth(c.Request().Header)
	if manager_id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	if err := h.middleware.AuthManager(manager_id); err != nil {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	code, accs := h.middleware.GetAccounts()
	return c.JSON(code, JSON{accs, ""})
}

func (h *ManagerHandler) FreezeAccount(c echo.Context) error {
	manager_id := h.middleware.Auth(c.Request().Header)
	if manager_id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	if err := h.middleware.AuthManager(manager_id); err != nil {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	id := t["Id"].(string)
	code, err := h.middleware.UpdateAccount(id, "FREEZED")
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	} else {
		return c.JSON(code, JSON{nil, ""})
	}
}

func (h *ManagerHandler) BlockAccount(c echo.Context) error {
	manager_id := h.middleware.Auth(c.Request().Header)
	if manager_id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	if err := h.middleware.AuthAdmin(manager_id); err != nil {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	id := t["Id"].(string)
	code, err := h.middleware.UpdateAccount(id, "BLOCKED")
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	} else {
		return c.JSON(code, JSON{nil, ""})
	}
}

func (h *ManagerHandler) GetTransactions(c echo.Context) error {
	manager_id := h.middleware.Auth(c.Request().Header)
	if manager_id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	if err := h.middleware.AuthManager(manager_id); err != nil {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	code, trs := h.middleware.GetTransactions()
	return c.JSON(code, JSON{trs, ""})
}

func (h *ManagerHandler) CancelTransaction(c echo.Context) error {
	manager_id := h.middleware.Auth(c.Request().Header)
	if manager_id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	if err := h.middleware.AuthManager(manager_id); err != nil {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	id := t["Id"].(string)

	code, err := h.middleware.CancelTransaction(id)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	} else {
		return c.JSON(code, JSON{nil, ""})
	}
}

func (h *ManagerHandler) GetUsers(c echo.Context) error {
	manager_id := h.middleware.Auth(c.Request().Header)
	if manager_id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	if err := h.middleware.AuthAdmin(manager_id); err != nil {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	code, usrs := h.middleware.GetUsers()
	return c.JSON(code, JSON{usrs, ""})
}

func (h *ManagerHandler) UpdateRole(c echo.Context) error {
	manager_id := h.middleware.Auth(c.Request().Header)
	if manager_id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}
	if err := h.middleware.AuthAdmin(manager_id); err != nil {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	id := t["Id"].(string)
	role := t["Role"].(string)
	code, err := h.middleware.UpdateRole(id, role)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	} else {
		return c.JSON(code, JSON{nil, ""})
	}
}
