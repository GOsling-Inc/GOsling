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

func (d *ManageDatabase) GetUnconfirmed() []models.Unconfirmed {
	var unc []models.Unconfirmed
	var ids []string

	query := "SELECT id FROM loans WHERE state = $1"
	d.db.Select(&ids, query, "PENDING")
	for _, id := range ids {
		unc = append(unc, models.Unconfirmed{
			Table: "loans",
			Id:    id,
		})
	}

	query = "SELECT id FROM deposits WHERE state = $1"
	d.db.Select(&ids, query, "PENDING")
	for _, id := range ids {
		unc = append(unc, models.Unconfirmed{
			Table: "deposits",
			Id:    id,
		})
	}

	query = "SELECT id FROM insurances WHERE state = $1"
	d.db.Select(&ids, query, "PENDING")
	for _, id := range ids {
		unc = append(unc, models.Unconfirmed{
			Table: "insurances",
			Id:    id,
		})
	}

	return unc
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

func (d *ManageDatabase) ConfirmDeposit(deposit models.Deposit) error {
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

func (d *ManageDatabase) GetUsers() []models.User {
	var users []models.User

	query := "SELECT * FROM users"
	d.db.Select(&users, query)

	for _, u := range users {
		u.Password = ""
	}
	return users
}

func (d *ManageDatabase) BlockUser(id string) error {
	query := "UPDATE users SET state = $1 WHERE id = #2"
	_, err := d.db.Exec(query, "BLOCKED", id)
	return err
}

func (d *ManageDatabase) GetTransactions() []models.Trasfer {
	var transfers []models.Trasfer

	query := "SELECT * FROM transactions"
	d.db.Select(&transfers, query)

	return transfers
}

func (d *AccountDatabase) CancelTransaction(trasaction models.Trasfer) error {
	query := "DELETE from transactions WHERE id = $1"
	d.db.Exec(query, trasaction.Id)

	return d.Transfer(trasaction.Receiver, trasaction.Sender, trasaction.Amount)
}

func (d *ManageDatabase) GetAccounts() []models.Account {
	var accounts []models.Account

	query := "SELECT * FROM accounts"
	d.db.Select(&accounts, query)

	return accounts
}

func (d *AccountDatabase) GetTransferById(id string) (models.Trasfer, error) {
	var transfer models.Trasfer
	query := "Select * from transfers WHERE id = $1"
	err := d.db.Get(&transfer, query, id)
	return transfer, err
}

func (d *ManageDatabase) UpdateAccount(id, state string) error {
	query := "UPDATE users SET state = $1 WHERE id = $2"
	_, err := d.db.Exec(query, state, id)
	return err
}

func (d *ManageDatabase) UpdateRole(id, role string) error {
	query := "UPDATE users SET role = $1 WHERE id = $2"
	_, err := d.db.Exec(query, role, id)
	return err
}
