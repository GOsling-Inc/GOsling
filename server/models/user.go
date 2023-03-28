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

type JWTClaims struct {
	jwt.StandardClaims
	ID string `json:"id"`
}
