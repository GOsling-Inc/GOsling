package services

import (
	"crypto/sha256"
	"errors"
	"time"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/golang-jwt/jwt"
)
/*
func (s *Service) SignIn() error {

	var input models.SignInInput
	//ref to DB...
	if input.Username == "Pavel" && input.Password == "1234" {
		cookie := &http.Cookie{}

		cookie.Name = "sessionId"
		cookie.Value = "value"
		cookie.Expires = time.Now().Add(48 * time.Hour)
		c.SetCookie(cookie)

		return c.String(http.StatusOK, "WelCUM")
	}
	return c.String(http.StatusUnauthorized, "Wrong username or password!")
}
*/
// func (s *Service) SignUp(c echo.Context) error
// {
// 	u := new(models.User)

// }

func (s *Service) HashPass(password string) (string, error) {
	if len(password) > 0 {
		hashed, err := s.getHashedPassword(password)
		if err != nil {
			return "", err
		}
		return hashed, nil
	}
	return "", errors.New("length of password must be more than 0")
}

func (s *Service) getHashedPassword(str string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}

	return string(str), nil
}

func CreateJWTToken(id string) (string, error) {
	claims := models.JWTClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
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