package services

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (s *Service) SignIn(c echo.Context) error {

	username := c.QueryParam("Pavel")
	password := c.QueryParam("1234")

	if username == "Pavel" && password == "1234" {
		cookie := &http.Cookie{}

		cookie.Name = "sessionId"
		cookie.Value = "value"
		cookie.Expires = time.Now().Add(48 * time.Hour)
		c.SetCookie(cookie)

		return c.String(http.StatusOK, "ABOBA")
	}
	return c.String(http.StatusUnauthorized, "Wrong username or password!")
}
