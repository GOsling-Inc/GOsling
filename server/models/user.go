package models

import (
	"crypto/sha256"
	"errors"
	"math/rand"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/golang-jwt/jwt"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type JWTClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
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
		hashed, err := u.getHashedPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = hashed
		return nil
	}
	return errors.New("length of password must be more than 0")
}

func (u *User) getHashedPassword(str string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}

	return string(str), nil
}
