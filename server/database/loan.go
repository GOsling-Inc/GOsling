package database

import (
	"context"
	"time"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/jmoiron/sqlx"
)

type ILoanDatabase interface {
	AddLoan(models.Loan) error
	GetUserLoans(string) ([]models.Loan, error)
	Debits() error
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
	ctx := context.Background()
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO loans (accountid, userid, amount, remaining, part, percent, period, deadline) values ($1, $2, $3, $4, $5, $6, $7, $8)", loan.AccountId, loan.UserId, loan.Amount, loan.Remaining, loan.Part, loan.Percent, loan.Period, loan.Deadline)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET amount = amount + $1 WHERE id = $2", loan.Amount, loan.AccountId)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func (d *LoanDatabase) GetUserLoans(userId string) ([]models.Loan, error) {
	var loans []models.Loan
	query := "SELECT * FROM loans WHERE userid=$1"
	err := d.db.Select(&loans, query, userId)
	return loans, err
}

func (d *LoanDatabase) Debits() error {
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
		_, err = tx.ExecContext(ctx, "UPDATE accounts SET amount = amount - $1 WHERE id = $2 AND status = 'ACTIVE'", loan.Part, loan.AccountId)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = tx.ExecContext(ctx, "UPDATE loans SET remaining = remaining - $1, deadline=deadline + INTERVAL '1 month' WHERE id = $2", loan.Part, loan.Id)
		if err != nil {
			tx.Rollback()
			return err
		}
		if loan.Remaining - loan.Part <= 0 {
			_, err = tx.ExecContext(ctx, "UPDATE loans SET state = 'CLOSED' WHERE id = $1", loan.Id)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		err = tx.Commit()
		return err
	}

	return nil
}
