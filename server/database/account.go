package database

import (
	"context"
	"fmt"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/jmoiron/sqlx"
)

type IAccountDatabase interface {
	AddAccount(account models.Account) error
	GetUserAccounts(userId string) ([]models.Account, error)
	GetAccountById(id string) (models.Account, error)
	DeleteAccount(id string) error
	UserTransfers(id string) []models.Trasfer
	Transfer(senderId, receiverId string, amount float64) error
	AddTransfer(transfer models.Trasfer) error
	CancelTransaction(trasaction models.Trasfer) error
	GetTransferById(id string) (models.Trasfer, error)
	Exchange(senderId, receiverId string, sender_amount, receiver_amount float64) error
	AddExchange(exchange models.Exchange) error
}

type AccountDatabase struct {
	db *sqlx.DB
}

func NewAccountDatabase(db *sqlx.DB) *AccountDatabase {
	return &AccountDatabase{
		db: db,
	}
}

func (d *AccountDatabase) AddAccount(account models.Account) error {
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
	query := "SELECT * FROM accounts WHERE id = $1"
	err := d.db.Get(&account, query, id)
	return account, err
}

func (d *AccountDatabase) DeleteAccount(id string) error {
	var empty string
	query := "UPDATE accounts SET state = 'DELETED' WHERE id=$1 RETURNING id"
	return d.db.Get(&empty, query, id)
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

func (d *AccountDatabase) AddTransfer(transfer models.Trasfer) error {
	var id string
	query := "INSERT INTO transfers (sender, receiver, amount) values ($1, $2, $3) RETURNING id"
	err := d.db.Get(&id, query, transfer.Sender, transfer.Receiver, transfer.Amount)
	return err
}

func (d *AccountDatabase) CancelTransaction(trasaction models.Trasfer) error {
	query := "DELETE from transactions WHERE id = $1"
	d.db.Exec(query, trasaction.Id)

	return d.Transfer(trasaction.Receiver, trasaction.Sender, trasaction.Amount)
}

func (d *AccountDatabase) UserTransfers(id string) []models.Trasfer {
	var transfer []models.Trasfer
	query := fmt.Sprintf("SELECT * from transfers WHERE receiver LIKE '%%%s%%' OR sender LIKE '%%%s%%'", id)
	err := d.db.Select(&transfer, query)
	fmt.Println(id, err, transfer)
	return transfer
}

func (d *AccountDatabase) GetTransferById(id string) (models.Trasfer, error) {
	var transfer models.Trasfer
	query := "SELECT * from transfers WHERE id = $1"
	err := d.db.Get(&transfer, query, id)
	return transfer, err
}

func (d *AccountDatabase) Exchange(senderId, receiverId string, sender_amount, receiver_amount float64) error {
	ctx := context.Background()
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET amount = amount + $1 WHERE id = $2", receiver_amount, receiverId)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET amount = amount - $1 WHERE id = $2", sender_amount, senderId)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func (d *AccountDatabase) AddExchange(exchange models.Exchange) error {
	var id string
	query := "INSERT INTO exchanges (sender, receiver, sender_amount, receiver_amount, course) values ($1, $2, $3, $4, $5) RETURNING id"
	err := d.db.Get(&id, query, exchange.Sender, exchange.Receiver, exchange.SenderAmount, exchange.ReceiverAmount, exchange.Course)
	return err
}
