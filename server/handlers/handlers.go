package handlers

import (
	"github.com/GOsling-Inc/GOsling/services"
)

type Handler struct {
	IAuthHandler
	IUserHandler
	IAccountHadler
	ILoanHandler
	IDepositHandler
}

func New(s *services.Service) *Handler {
	return &Handler{
		IAuthHandler:    NewAuthHandler(s),
		IUserHandler:    NewUserHandler(s),
		IAccountHadler:  NewAccountHandler(s),
		ILoanHandler:    NewLoanHandler(s),
		IDepositHandler: NewDepositHandler(s),
	}
}
