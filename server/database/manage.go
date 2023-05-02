package database

import (
	"context"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/jmoiron/sqlx"
)

type ManageDatabase struct {
	db *sqlx.DB
}

func NewManageDatabase(db *sqlx.DB) *ManageDatabase {
	return &ManageDatabase{
		db: db,
	}
}

func (d *ManageDatabase) ConfirmLoan(loan models.Loan) error {
	ctx := context.Background()
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, "UPDATE loans SET state = $1 WHERE id = $2", loan.State, loan.Id)
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

func (d *ManageDatabase) ConfirmInsurance(insurance models.Insurance) error {
	query := "UPDATE accounts SET amount = amount - $1 WHERE id = $3"
	_, err := d.db.Exec(query, insurance.Amount, insurance.UserId)
	return err
}

func (d *DepositDatabase) ConfirmDeposit(deposit models.Deposit) error {
	ctx := context.Background()
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, "UPDATE deposits SET state = $1 WHERE id = $2", deposit.State, deposit.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET amount = amount - $1 WHERE id = $2", deposit.Amount, deposit.AccountId)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}
