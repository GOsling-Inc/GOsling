package services

import (
	"github.com/GOsling-Inc/GOsling/database"
)

type IService interface {
	IAuthService
	IUserService
	IAccountService
	ILoanService
	IDepositService
	IInsuranceService
	IInvestmentService
	IManageService
}

type Service struct {
	IAuthService
	IUserService
	IAccountService
	ILoanService
	IDepositService
	IInsuranceService
	IInvestmentService
	IManageService
}

func New(d database.IDatabase) *Service {
	return &Service{
		IAuthService:       NewAuthService(d),
		IUserService:       NewUserService(d),
		IAccountService:    NewAccountService(d),
		ILoanService:       NewLoanService(d),
		IDepositService:    NewDepositService(d),
		IInsuranceService:  NewInsuranceService(d),
		IInvestmentService: NewInvestmentService(d),
		IManageService:     NewManageService(d),
	}
}
