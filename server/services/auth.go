package services

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

type SignInInput struct {
	username string
	password string
}

type JWTClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func CreateJWTToken() (string, error) {
	claims := JWTClaims{
		"userid",
		jwt.StandardClaims{
			Id:        "user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte("a"))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *Service) SignIn(c echo.Context) error {

	var input SignInInput
	//ref to DB...
	if input.username == "Pavel" && input.password == "1234" {
		cookie := &http.Cookie{}

		cookie.Name = "sessionId"
		cookie.Value = "value"
		cookie.Expires = time.Now().Add(48 * time.Hour)
		c.SetCookie(cookie)

		return c.String(http.StatusOK, "ABOBA")
	}
	return c.String(http.StatusUnauthorized, "Wrong username or password!")
}
