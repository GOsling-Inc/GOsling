package middleware

import (
	"errors"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
)

type ILoanMiddleware interface {
	ProvideLoan(models.Loan) (int, error)
	GetUserLoans(string) (int, []models.Loan, error)
}

type LoanMiddleware struct {
	service services.IService
}

func NewLoanMiddleware(s services.IService) *LoanMiddleware {
	return &LoanMiddleware{
		service: s,
	}
}

func (m *LoanMiddleware) ProvideLoan(loan models.Loan) (int, error) {

	acc, err := m.service.GetAccountById(loan.AccountId)
	if err != nil || acc.UserId != loan.UserId {
		return UNAUTHORIZED, errors.New("invalid account")
	}
	if err = m.service.ProvideLoan(loan); err != nil {
		return INTERNAL, err
	}
	return ACCEPTED, nil
}

func (m *LoanMiddleware) GetUserLoans(userId string) (int, []models.Loan, error) {
	loans, err := m.service.GetUserLoans(userId)
	if err != nil {
		return INTERNAL, nil, err
	}
	return OK, loans, nil
}
