package models

import (
	"crypto/sha256"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

var (
	hashed string
	err    error
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func (u *User) Init(name string, surname string, email string, password string) User {
	return User{
		Name:     name,
		Surname:  surname,
		Email:    email,
		Password: password,
	}
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 100)))
}

func (u *User) HashPass() error {
	if len(u.Password) > 0 {
		hashed, err = getHashedPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = hashed
	}
	return nil
}

func getHashedPassword(s string) (string, error) {
	hash := sha256.New()
	_, err = hash.Write([]byte(s))
	if err != nil {
		return "", err
	}

	return string(s), nil
}
