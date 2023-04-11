package handlers

import (
	"github.com/GOsling-Inc/GOsling/middleware"
)

type JSON struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

type OBJ map[string]interface{}

type Handler struct {
	*AuthHandler
	*UserHandler
	*AccountHandler
	*LoanHandler
	*DepositHandler
	*EnsuranceHandler
}

func New(m *middleware.Middleware) *Handler {
	return &Handler{
		AuthHandler:      NewAuthHandler(m),
		UserHandler:      NewUserHandler(m),
		AccountHandler:   NewAccountHandler(m),
		LoanHandler:      NewLoanHandler(m),
		DepositHandler:   NewDepositHandler(m),
		EnsuranceHandler: NewEnsuranceHandler(m),
	}
}
