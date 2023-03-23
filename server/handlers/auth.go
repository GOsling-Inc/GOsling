package handlers

import (
	"net/http"
	"time"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
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
	if err := h.service.SignUp(user); err != nil {
		c.JSON(401, err.Error())
		return err
	}
	token, err := h.CreateJWT(user.ID)
	if err != nil {
		c.JSON(401, err.Error())
		return err
	}
	c.SetCookie(h.CreateCookie(token))
	return c.Redirect(200, "/")
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
	if err := h.service.SignIn(user); err != nil {
		c.JSON(401, err.Error())
		return err
	}
	token, err := h.CreateJWT(user.ID)
	if err != nil {
		c.JSON(401, err.Error())
		return err
	}
	c.SetCookie(h.CreateCookie(token))
	return c.Redirect(200, "/")
}

func (h *Handler) CreateCookie(token string) *http.Cookie {
	cookie := &http.Cookie{}
	cookie.Name = "Token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	return cookie
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
	token, err := rawToken.SignedString([]byte("a"))
	if err != nil {
		return "", err
	}
	return token, nil
}

/*
cookie := &http.Cookie{}

		cookie.Name = "sessionId"
		cookie.Value = "value"
		cookie.Expires = time.Now().Add(48 * time.Hour)
		c.SetCookie(cookie)

		return c.String(http.StatusOK, "WelCUM")
	}
	return c.String(http.StatusUnauthorized, "Wrong username or password!")
*/
