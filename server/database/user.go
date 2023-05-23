package database

import (
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/jmoiron/sqlx"
)

type IUserDatabase interface {
	GetUserByMail(string) (models.User, error)
	GetUserById(string) (models.User, error)
	AddUser(*models.User) error
	UpdatePasswordUser(string, string) error
	UpdateUserData(string, string, string, string) error
}

type UserDatabase struct {
	db *sqlx.DB
}

func NewUserDatabase(db *sqlx.DB) *UserDatabase {
	return &UserDatabase{
		db: db,
	}
}

func (d *UserDatabase) GetUserByMail(mail string) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE email=$1"
	err := d.db.Get(&user, query, mail)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (d *UserDatabase) GetUserById(id string) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE id=$1"
	err := d.db.Get(&user, query, id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (d *UserDatabase) AddUser(user *models.User) error {
	var id string
	query := "INSERT INTO users (id, name, surname, email, password, birthdate) values ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := d.db.Get(&id, query, user.Id, user.Name, user.Surname, user.Email, user.Password, user.Birthdate)
	user.Id = id
	return err
}

func (d *UserDatabase) UpdatePasswordUser(id, password string) error {
	var ID string
	query := "UPDATE users SET password=$1 WHERE id=$2 RETURNING id"
	err := d.db.Get(&ID, query, password, id)
	return err
}

func (d *UserDatabase) UpdateUserData(id, name, surname, birthdate string) error {
	var ID string
	query := "UPDATE users SET name=$1, surname=$2, birthdate=$3 WHERE id=$4 RETURNING id"
	err := d.db.Get(&ID, query, name, surname, birthdate, id)
	return err
}
