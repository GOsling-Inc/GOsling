package models

import (
	"crypto/sha256"
	"math/rand"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

var (
	hashed string
	err    error
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func (u *User) FormID() {
	var charset = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, 7)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	u.ID = string(b)
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
