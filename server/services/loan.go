package services

import (
	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/utils"
)

type ILoantService interface {
	ProvideLoan(models.Loan) error
	GetUserLoans(string) ([]models.Loan, error)
}

type LoanService struct {
	database *database.Database
	Utils    *utils.Utils
}

func NewLoanService(d *database.Database, u *utils.Utils) *LoanService {
	return &LoanService{
		database: d,
		Utils:    u,
	}
}

func (s *LoanService) ProvideLoan(models.Loan) error {
	return nil
}

func (s *LoanService) GetUserLoans(string) ([]models.Loan, error) {
	return nil, nil
}