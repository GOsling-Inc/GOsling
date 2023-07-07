package middleware

import (
	"github.com/GOsling-Inc/GOsling/services"
)

const (
	OK           = 200
	CREATED      = 201
	ACCEPTED     = 202
	UNAUTHORIZED = 401
	FORBIDDEN    = 403
	INTERNAL     = 500
)

type IMiddleware interface {
	IAuthMiddleware
	IUserMiddleware
	IAccountMiddleware
	ILoanMiddleware
	IDepositMiddleware
	IInsuranceMiddleware
	IInvestmentMiddleware
	IManagerMiddleware
}

type Middleware struct {
	IAuthMiddleware
	IUserMiddleware
	IAccountMiddleware
	ILoanMiddleware
	IDepositMiddleware
	IInsuranceMiddleware
	IInvestmentMiddleware
	IManagerMiddleware
}

func New(s services.IService) *Middleware {
	return &Middleware{
		IAuthMiddleware:       NewAuthMiddleware(s),
		IUserMiddleware:       NewUserMiddleware(s),
		IAccountMiddleware:    NewAccountMiddleware(s),
		ILoanMiddleware:       NewLoanMiddleware(s),
		IDepositMiddleware:    NewDepositMiddleware(s),
		IInsuranceMiddleware:  NewInsuranceMiddleware(s),
		IInvestmentMiddleware: NewInvestmentMiddleware(s),
		IManagerMiddleware:    NewManagerMiddleware(s),
	}
}
