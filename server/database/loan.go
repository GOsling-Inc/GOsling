package database

import (
	"context"
	"time"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/jmoiron/sqlx"
)

type LoanDatabase struct {
	db *sqlx.DB
}

func NewLoanDatabase(db *sqlx.DB) *LoanDatabase {
	return &LoanDatabase{
		db: db,
	}
}

func (d *InsuranceDatabase) AddLoan(loan models.Loan) error {
	var id string
	query := "INSERT INTO loans (accountid, userid, amount, remaining, part, percent, period, deadline) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	return d.db.Get(&id, query, loan.AccountId, loan.UserId, loan.Amount, loan.Remaining, loan.Part, loan.Percent, loan.Period, loan.Deadline)
}

func (d *InsuranceDatabase) GetLoanById(id string) (models.Loan, error) {
	var loan models.Loan
	query := "SELECT * FROM loans WHERE userid=$1"
	err := d.db.Get(&loan, query, id)
	return loan, err
}

func (d *LoanDatabase) GetUserLoans(userId string) ([]models.Loan, error) {
	var loans []models.Loan
	query := "SELECT * FROM loans WHERE userid=$1"
	err := d.db.Select(&loans, query, userId)
	return loans, err
}

func (d *LoanDatabase) UpdateLoans() error {
	date := time.Now().Format("2006-01-02")
	var loans []models.Loan
	query := "SELECT * FROM loans WHERE deadline=$1"
	d.db.Select(&loans, query, date)
	ctx := context.Background()
	for _, loan := range loans {
		tx, err := d.db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(ctx, "UPDATE accounts SET amount = amount - $1 WHERE id = $2 AND state = 'ACTIVE'", loan.Part, loan.AccountId)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = tx.ExecContext(ctx, "UPDATE loans SET remaining = remaining - $1, deadline=deadline + INTERVAL '1 month' WHERE id = $2", loan.Part, loan.Id)
		if err != nil {
			tx.Rollback()
			return err
		}
		if loan.Remaining-loan.Part <= 0 {
			_, err = tx.ExecContext(ctx, "UPDATE loans SET state = 'CLOSED' WHERE id = $1", loan.Id)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
	}

	return nil
}
