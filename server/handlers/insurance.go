package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/labstack/echo/v4"
)

type IInsuranceHandler interface {
	POST_NewInsurance(echo.Context) error
	GET_User_Insurances(echo.Context) error
}

type InsuranceHandler struct {
	middleware middleware.IMiddleware
}

func NewInsuranceHandler(m middleware.IMiddleware) *InsuranceHandler {
	return &InsuranceHandler{
		middleware: m,
	}
}

func (h *InsuranceHandler) POST_NewInsurance(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	insurance := models.Insurance{
		UserId:    id,
		AccountId: t["AccountId"].(string),
		Period:    t["Period"].(string),
	}
	insurance.Amount, _ = strconv.ParseFloat(t["Amount"].(string), 64)
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
