package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/golang-jwt/jwt"
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

type SignInInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type JWTClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 100)))
}

