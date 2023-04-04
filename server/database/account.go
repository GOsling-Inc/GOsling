package database

import (
	"context"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/jmoiron/sqlx"
)

type IAccountDatabase interface {
	AddAccount(*models.Account) error
	GetUserAccounts(userId string) ([]models.Account, error)
	GetAccountById(id string) (models.Account, error)
	Transfer(senderId, receiverId string, amount float64) error
	AddTransfer(transfer *models.Trasfer) error
}

type AccountDatabase struct {
	db *sqlx.DB
}

func NewAccountDatabase(db *sqlx.DB) *AccountDatabase {
	return &AccountDatabase{
		db: db,
	}
}

func (d *AccountDatabase) AddAccount(account *models.Account) error {
	var id string
	query := "INSERT INTO accounts (id, userid, name, type, unit) values ($1, $2, $3, $4, $5) RETURNING id"
	err := d.db.Get(&id, query, account.Id, account.UserId, account.Name, account.Type, account.Unit)
	return err
}

func (d *AccountDatabase) GetUserAccounts(userId string) ([]models.Account, error) {
	var accounts []models.Account
	query := "SELECT * FROM accounts WHERE userId=$1"
	err := d.db.Select(&accounts, query, userId)
	return accounts, err
}

func (d *AccountDatabase) GetAccountById(id string) (models.Account, error) {
	var account models.Account
	query := "SELECT * FROM accounts WHERE id=$1"
	err := d.db.Get(&account, query, id)
	return account, err
}

func (d *AccountDatabase) Transfer(senderId, receiverId string, amount float64) error {
	ctx := context.Background()
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET amount = amount + $1 WHERE id = $2", amount, receiverId)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET amount = amount - $1 WHERE id = $2", amount, senderId)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func (d *AccountDatabase) AddTransfer(transfer *models.Trasfer) error {
	var id string
	query := "INSERT INTO transfers (sender, receiver, amount) values ($1, $2, $3) RETURNING id"
	err := d.db.Get(&id, query, transfer.Sender, transfer.Receiver, transfer.Amount)
	return err
}
