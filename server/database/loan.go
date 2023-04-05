package database

import (
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/jmoiron/sqlx"
)

type ILoanDatabase interface {
	AddLoan(models.Loan) error
	GetUserLoans(string) ([]models.Loan, error)
}

type LoanDatabase struct {
	db *sqlx.DB
}

func NewLoanDatabase(db *sqlx.DB) *LoanDatabase {
	return &LoanDatabase{
		db: db,
	}
}

func (d *LoanDatabase) AddLoan(loan models.Loan) error {
	var id string
	query := "INSERT INTO loans (accountid, userid, amount, remaining, part, percent, period, deadline) values ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	err := d.db.Get(&id, query, loan.AccountId, loan.UserId, loan.Amount, loan.Remaining, loan.Part, loan.Percent, loan.Period, loan.Deadline)
	return err
}

func (d *LoanDatabase) GetUserLoans(userId string) ([]models.Loan, error) {
	var loans []models.Loan
	query := "SELECT * FROM loans WHERE userid=$1"
	err := d.db.Select(&loans, query, userId)
	return loans, err
}
