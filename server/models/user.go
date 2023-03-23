package models

import (
<<<<<<< HEAD
=======
	"crypto/sha256"
	"math/rand"

>>>>>>> 21875465c2afd7cf8fee78adc7edacb5c4e02b7c
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

<<<<<<< HEAD
type SignInInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type JWTClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
=======
func (u *User) FormID() {
	var charset = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, 7)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	u.ID = string(b)
>>>>>>> 21875465c2afd7cf8fee78adc7edacb5c4e02b7c
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 100)))
}

