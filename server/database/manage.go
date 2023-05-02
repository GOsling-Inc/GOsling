package database

import (
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
	query := "UPDATE accounts SET amount = amount + $1, state = $2 WHERE id = $3"
	_, err := d.db.Exec(query, loan.Amount, loan.State, loan.Id)
	return err
}

func (d *ManageDatabase) ConfirmInsurance(deposit models.Insurance) error {
	query := "UPDATE accounts SET amount = amount - $1, state = $2 WHERE id = $3"
	_, err := d.db.Exec(query, deposit.Amount, deposit.State, deposit.Id)
	return err
}

func (d *DepositDatabase) ConfirmDeposit(deposit models.Deposit) error {
	query := "UPDATE accounts SET amount = amount - $1, state = $2 WHERE id = $3"
	_, err := d.db.Exec(query, deposit.Amount, deposit.State, deposit.Id)
	return err
}
