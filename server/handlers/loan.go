package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/labstack/echo/v4"
)

type ILoanHandler interface {
	POST_Loan(echo.Context) error
	GET_User_Loans(echo.Context) error
}

type LoanHandler struct {
	middleware middleware.IMiddleware
}

func NewLoanHandler(s middleware.IMiddleware) *LoanHandler {
	return &LoanHandler{
		middleware: s,
	}
}

func (h *LoanHandler) POST_Loan(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	decoder := json.NewDecoder(c.Request().Body)
	var t map[string]interface{}
	decoder.Decode(&t)

	loan := models.Loan{
		UserId:    id,
		AccountId: t["AccountId"].(string),
		Period:    t["Period"].(string),
	}
	loan.Amount, _ = strconv.ParseFloat(t["Amount"].(string), 64)
	loan.Percent, _ = strconv.ParseFloat(t["Percent"].(string), 64)

	code, err := h.middleware.ProvideLoan(loan)
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{"ok", ""})
}

func (h *LoanHandler) GET_User_Loans(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	code, loans, err := h.middleware.GetUserLoans(id)
	for i := 0; i < len(loans); i++ {
		loans[i].Period = loans[i].Period[0:10]
		loans[i].Deadline = loans[i].Deadline[0:10]
	}
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{loans, ""})
}
