package models

import "github.com/golang-jwt/jwt"

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	Role      string `json:"role"`
	Birthdate string `json:"date"`
}

type JWTClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
