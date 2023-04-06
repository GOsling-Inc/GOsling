package handlers

import (
	"strconv"
	"time"

	"github.com/GOsling-Inc/GOsling/models"
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

func (h *LoanHandler) POST_Loan(c echo.Context) error {
	beta_loan := &models.Loan{
		AccountId: c.FormValue("AccountId"),
		Period:    c.FormValue("Period"),
	}
	beta_loan.Amount, _ = strconv.ParseFloat(c.FormValue("Amount"), 64)
	beta_loan.Percent, _ = strconv.ParseFloat(c.FormValue("Percent"), 64)
	per, _ := strconv.Atoi(beta_loan.Period)
	beta_loan.Period = time.Now().AddDate(per, 0, 0).Format("2006-01-02")
	beta_loan.Deadline = time.Now().AddDate(0, 30, 0).Format("2006-01-02")
	header := c.Request().Header
	id, err := h.service.ParseJWT(header["Token"][0])
	if err != nil {
		return c.JSON(401, err.Error())
	}
	_, err = h.service.GetUser(id)
	if err != nil {
		return c.JSON(401, err.Error())
	}
	beta_loan.UserId = id
	beta_loan.Remaining = beta_loan.Amount + beta_loan.Amount*beta_loan.Percent/100
	if err = h.service.ProvideLoan(beta_loan); err != nil {
		return c.JSON(401, err.Error())
	}
	return nil
}

func (h *LoanHandler) GET_User_Loans(c echo.Context) error {
	header := c.Request().Header
	id, err := h.service.ParseJWT(header["Token"][0])
	if err != nil {
		return c.JSON(401, err.Error())
	}
	_, err = h.service.GetUser(id)
	if err != nil {
		return c.JSON(401, err.Error())
	}
	loans, err := h.service.GetUserLoans(id)
	if err != nil {
		return c.JSON(401, err.Error())
	}
	return c.JSON(200, loans)
}
