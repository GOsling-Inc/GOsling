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

type Middleware struct {
	*AuthMiddleware
	*UserMiddleware
	*AccountMiddleware
	*LoanMiddleware
	*DepositMiddleware
	*InsuranceMiddleware
	*InvestmentMiddleware
}

func New(s *services.Service) *Middleware {
	return &Middleware{
		AuthMiddleware:       NewAuthMiddleware(s),
		UserMiddleware:       NewUserMiddleware(s),
		AccountMiddleware:    NewAccountMiddleware(s),
		LoanMiddleware:       NewLoanMiddleware(s),
		DepositMiddleware:    NewDepositMiddleware(s),
		InsuranceMiddleware:  NewInsuranceMiddleware(s),
		InvestmentMiddleware: NewInvestmentMiddleware(s),
	}
}
