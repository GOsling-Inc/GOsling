package services_test //debug needed

import (
	"net/http"
	"testing"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/handlers"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	db   database.Database
	hand handlers.AuthHandler
)

func Test_User_SignIn(t *testing.T) {

	server := echo.New()
	user := models.User{
		Name:      "ABOBA",
		Surname:   "ABOBOB",
		Email:     "examle@gmail.com",
		Password:  "12345678",
		Role:      "user",
		Birthdate: "2020-01-01",
	}
	db.AddUser(&user)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    user.Email,
				"password": user.Password,
			},
			expectedCode: http.StatusAccepted,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tmp := server.POST("/sign-in", hand.POST_SignIn)
			assert.Equal(t, tc.expectedCode, tmp)
		})
	}
}

func Test_User_SignUp(t *testing.T) {
	server := echo.New()
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"name":      "ABOBA",
				"surname":   "ABOBOB",
				"email":     "examle@gmail.com",
				"password":  "12345678",
				"role":      "user",
				"birthdate": "2020-01-01",
			},
			expectedCode: http.StatusAccepted,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tmp := server.POST("/sign-in", hand.POST_SignUp)
			assert.Equal(t, tc.expectedCode, tmp)
		})
	}
}
