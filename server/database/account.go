package database

import (
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/jmoiron/sqlx"
)

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
