package handlers

import (
	"strconv"

	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/labstack/echo/v4"
)

type LoanHandler struct {
	middleware *middleware.Middleware
}

func NewLoanHandler(s *middleware.Middleware) *LoanHandler {
	return &LoanHandler{
		middleware: s,
	}
}

func (h *LoanHandler) POST_Loan(c echo.Context) error {
	id := h.middleware.Auth(c.Request().Header)
	if id == "" {
		return c.JSON(middleware.UNAUTHORIZED, JSON{nil, "invalid token"})
	}

	loan := models.Loan{
		UserId: id,
		AccountId: c.FormValue("AccountId"),
		Period:    c.FormValue("Period"),
	}
	loan.Amount, _ = strconv.ParseFloat(c.FormValue("Amount"), 64)
	loan.Percent, _ = strconv.ParseFloat(c.FormValue("Percent"), 64)

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
	if err != nil {
		return c.JSON(code, JSON{nil, err.Error()})
	}
	return c.JSON(code, JSON{loans, ""})
}
