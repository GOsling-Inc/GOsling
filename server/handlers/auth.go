package handlers

import (
	"time"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const (
	salt = "GOsling"
)

type AuthHandler struct {
	service *services.Service
}

func NewAuthHandler(s *services.Service) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

func (h *AuthHandler) POST_SignUp(c echo.Context) error {
	user := models.User{
		Id:       h.service.MakeID(),
		Name:     c.FormValue("Name"),
		Surname:  c.FormValue("Surname"),
		Email:    c.FormValue("Email"),
		Password: c.FormValue("Password"),
		Role:     "user",
	}
	if err := h.service.Validate(&user); err != nil {
		c.JSON(401, err.Error())
		return err
	}
	if err := h.service.HashPassword(&user); err != nil {
		c.JSON(401, err.Error())
		return err
	}
	if err := h.service.SignUp(&user); err != nil {
		c.JSON(401, err.Error())
		return err
	}
	token, err := h.CreateJWT(user.Id)
	if err != nil {
		c.JSON(401, err.Error())
		return err
	}
	return c.JSON(201, map[string]string{
		"Token": token,
	})
}

func (h *AuthHandler) POST_SignIn(c echo.Context) error {
	user := models.User{
		Email:    c.FormValue("Email"),
		Password: c.FormValue("Password"),
	}
	if err := h.service.Validate(&user); err != nil {
		c.JSON(401, err.Error())
		return err
	}
	if err := h.service.HashPassword(&user); err != nil {
		c.JSON(401, err.Error())
		return err
	}
	if err := h.service.SignIn(&user); err != nil {
		c.JSON(401, err.Error())
		return err
	}
	token, err := h.CreateJWT(user.Id)
	if err != nil {
		c.JSON(401, err.Error())
		return err
	}
	return c.JSON(201, map[string]string{
		"Token": token,
	})
}

type JWTClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func (h *AuthHandler) CreateJWT(id string) (string, error) {
	claims := JWTClaims{
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
