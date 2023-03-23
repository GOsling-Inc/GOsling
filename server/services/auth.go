package services

import (
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