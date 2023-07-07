package handlers

import (
	"github.com/GOsling-Inc/GOsling/middleware"
)

type JSON struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

type OBJ map[string]interface{}

type IHandler interface {
	IAuthHandler
	IUserHandler
	IAccountHandler
	ILoanHandler
	IDepositHandler
	IInsuranceHandler
	IInvestmentHandler
	IManagerHandler
}

type Handler struct {
	IAuthHandler
	IUserHandler
	IAccountHandler
	ILoanHandler
	IDepositHandler
	IInsuranceHandler
	IInvestmentHandler
	IManagerHandler
}

func New(m middleware.IMiddleware) *Handler {
	return &Handler{
		IAuthHandler:       NewAuthHandler(m),
		IUserHandler:       NewUserHandler(m),
		IAccountHandler:    NewAccountHandler(m),
		ILoanHandler:       NewLoanHandler(m),
		IDepositHandler:    NewDepositHandler(m),
		IInsuranceHandler:  NewInsuranceHandler(m),
		IInvestmentHandler: NewInvestmentHandler(m),
		IManagerHandler:    NewManagerHandler(m),
	}
}
