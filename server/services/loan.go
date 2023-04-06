package services

import (
	"errors"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/utils"
)

type ILoantService interface {
	ProvideLoan(*models.Loan) error
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

func (s *LoanService) ProvideLoan(loan *models.Loan) error {
	loans, err := s.GetUserLoans(loan.UserId)
	if err != nil {
		return err
	}
	for _, t := range loans {
		if t.State == "ACTIVE" {
			return errors.New("can't open a new loan")
		}
	}
	if err = s.database.Debits(); err != nil {
		return err
	}
	err = s.database.AddLoan(*loan)
	return err
}

func (s *LoanService) GetUserLoans(id string) ([]models.Loan, error) {
	loans, err := s.database.GetUserLoans(id)
	return loans, err
}
