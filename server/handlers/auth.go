package handlers

import (
	"time"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const (
	salt = "GOsling"
)

func (h *Handler) POST_SignUp(c echo.Context) error {
	user := models.User{
		Name:     c.FormValue("Name"),
		Surname:  c.FormValue("Surname"),
		Email:    c.FormValue("Email"),
		Password: c.FormValue("Password"),
	}
	if err := user.Validate(); err != nil {
		c.JSON(401, err.Error())
		return err
	}
	if err := user.HashPass(); err != nil {
		c.JSON(401, err.Error())
		return err
	}
	if err := h.service.SignUp(&user); err != nil {
		c.JSON(401, err.Error())
		return err
	}
	token, err := h.CreateJWT(user.ID)
	if err != nil {
		c.JSON(401, err.Error())
		return err
	}
	return c.JSON(201, map[string]string{
		"Token": token,
	})
}

func (h *Handler) POST_SignIn(c echo.Context) error {
	user := models.User{
		Email:    c.FormValue("Email"),
		Password: c.FormValue("Password"),
	}
	if err := user.Validate(); err != nil {
		c.JSON(401, err.Error())
		return err
	}
	if err := user.HashPass(); err != nil {
		c.JSON(401, err.Error())
		return err
	}
	if err := h.service.SignIn(&user); err != nil {
		c.JSON(401, err.Error())
		return err
	}
	token, err := h.CreateJWT(user.ID)
	if err != nil {
		c.JSON(401, err.Error())
		return err
	}
	return c.JSON(201, map[string]string{
		"Token": token,
	})
}

func (h *Handler) CreateJWT(id string) (string, error) {
	claims := models.JWTClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			Id:        "user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(salt))
	if err != nil {
		return "", err
	}
	return token, nil
}