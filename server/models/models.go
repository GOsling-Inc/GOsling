package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	Role      string `json:"role"`
	Birthdate string `json:"birthdate"`
}

type Account struct {
	Id     string  `json:"id"`
	UserId string  `json:"userid"`
	Name   string  `json:"name"`
	Type   string  `json:"type"` // BASIC / BUSINESS / INVESTMENT
	Unit   string  `json:"unit"` // BYN / USD / EUR
	Amount float32 `json:"amount"`
	State  string  `json:"state"` // ACTIVE / BLOCKED / FROZEN
}

type Trasfer struct {
	Id     string  `json:"id"`
	Receiver string  `json:"receiver"`
	Sender string  `json:"sender"`
	Amount float32 `json:"amount"`
}

type JWTClaims struct {
	jwt.StandardClaims
	ID string `json:"id"`
}
