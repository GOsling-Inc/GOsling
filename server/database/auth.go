package database

import (
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/jmoiron/sqlx"
)

type AuthDatabase struct {
	db *sqlx.DB
}

func NewAuthDatabase(db *sqlx.DB) *AuthDatabase {
	return &AuthDatabase{
		db: db,
	}
}

func (d *AuthDatabase) GetUserByMail(mail string) (*models.User, error) {
	user := new(models.User)

	query := "SELECT * FROM users WHERE email=$1"
	err := d.db.Get(user, query, mail)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *AuthDatabase) GetUserById(id string) (*models.User, error) {
	user := new(models.User)

	query := "SELECT * FROM users WHERE id=$1"
	err := d.db.Get(user, query, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *AuthDatabase) AddUser(user *models.User) error {
	var id string
	query := "INSERT INTO users (id, name, surname, email, password) values ($1, $2, $3, $4, $5) RETURNING id"
	err := d.db.Get(&id, query, user.Id, user.Name, user.Surname, user.Email, user.Password)
	return err
}
