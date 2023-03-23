package handlers

import (
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) POST_SignUp(c echo.Context) error {
	user := models.User{
		Name: c.FormValue("Name"),
		Surname: c.FormValue("Surname"),
		Email: c.FormValue("Emaul"),
		Password: c.FormValue("Password"),
	}
	if err := user.Validate(); err != nil {
		c.JSON(401, err.Error())
		return err
	}
	hashed, err := h.service.HashPass(user.Password)
	if err != nil {
		c.JSON(401, err.Error())
		return err
	}
	user.Password = hashed

}

func (h *Handler) POST_SignIn(c echo.Context) error {
	user := models.SignInInput{
		Email: c.FormValue("Emaul"),
		Password: c.FormValue("Password"),
	}

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