package handlers

import (
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/labstack/echo/v4"
)

type ILoanHandler interface {
	POST_Loan(echo.Context) error
	GET_User_Loans(echo.Context) error
}

type LoanHandler struct {
	service *services.Service
}

func NewLoanHandler(s *services.Service) *LoanHandler {
	return &LoanHandler{
		service: s,
	}
}

func (h * LoanHandler) POST_Loan(c echo.Context) error {
	return nil
}

func (h * LoanHandler) GET_User_Loans(c echo.Context) error {
	return nil
}