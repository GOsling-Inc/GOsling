package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
)

type LoanService struct {
	database *database.Database
}

func NewLoanService(d *database.Database) *LoanService {
	return &LoanService{
		database: d,
	}
}

func (s *LoanService) ProvideLoan(loan models.Loan) error {
	loans, err := s.GetUserLoans(loan.UserId)
	if err != nil {
		return err
	}
	for _, t := range loans {
		if t.State == "ACTIVE" {
			return errors.New("can't open a new loan")
		}
	}
	per, _ := strconv.Atoi(loan.Period)
	loan.Period = time.Now().AddDate(per, 0, 0).Format("2006-01-02")
	loan.Deadline = time.Now().AddDate(0, 0, 30).Format("2006-01-02")
	loan.Remaining = loan.Amount + loan.Amount * loan.Percent / 100 * float64(per)
	loan.Part = loan.Remaining / (12 * float64(per))
	err = s.database.AddLoan(loan)
	return err
}

func (s *LoanService) GetUserLoans(id string) ([]models.Loan, error) {
	return s.database.GetUserLoans(id)
}
