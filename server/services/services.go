package services

import (
	"github.com/GOsling-Inc/GOsling/database"
)

type Service struct {
	*AuthService
	*UserService
	*AccountService
	*LoanService
	*DepositService
	*InsuranceService
	*InvestmentService
}

func New(d *database.Database) *Service {
	return &Service{
		AuthService:       NewAuthService(d),
		UserService:       NewUserService(d),
		AccountService:    NewAccountService(d),
		LoanService:       NewLoanService(d),
		DepositService:    NewDepositService(d),
		InsuranceService:  NewInsuranceService(d),
		InvestmentService: NewInvestmentService(d),
	}
}
