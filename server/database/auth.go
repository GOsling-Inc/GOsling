package database

import (
	"log"

	"github.com/GOsling-Inc/GOsling/models"
)

func (d *Database) GetUserByMail(mail string) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE email=$1"
	err := d.db.Get(&user, query, mail)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (d *Database) GetUserById(id string) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE id=$1"
	err := d.db.Get(&user, query, id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (d *Database) AddUser(user *models.User) error {
	var id string
	log.Println(user)
	query := "INSERT INTO users (id, name, surname, email, password) values ($1, $2, $3, $4, $5) RETURNING id"
	err := d.db.Get(&id, query, user.ID, user.Name, user.Surname, user.Email, user.Password)
	return err
}