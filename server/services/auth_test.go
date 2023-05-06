package services_test //debug needed

import (
	"errors"
	"testing"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/stretchr/testify/assert"
)

var (
	db database.MockDatabase
	//serv  services.AuthService
	user0 = models.User{
		Id:        "1234",
		Name:      "User",
		Surname:   "Existing",
		Email:     "imhere@mail.com",
		Password:  "secRetec0de",
		Role:      "user",
		Birthdate: "1979-09-21",
	}
)

func Test_Auth_SignUp(t *testing.T) {
	db = *database.NewMock()
	//serv = *services.NewAuthService(db)
	db.AddUser(&user0)
	testCases := []struct {
		name        string
		payload     models.User
		expectedErr string
	}{
		{
			name: "valid",
			payload: models.User{
				Id:        "",
				Name:      "ABOBA",
				Surname:   "ABOBOV",
				Email:     "test1@gmail.com",
				Password:  "154492Ad",
				Role:      "user",
				Birthdate: "2002-01-01",
			},
			expectedErr: "",
		},
		{
			name: "user_exists",
			payload: models.User{
				Id:        "",
				Name:      "Same",
				Surname:   "Email",
				Email:     "imhere@mail.com",
				Password:  "154492Ad",
				Role:      "user",
				Birthdate: "2002-01-01",
			},
			expectedErr: "user with this email already registered",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}

func Test_Auth_SignIn(t *testing.T) {
	db = *database.NewMock()
	db.AddUser(&user0)
	testCases := []struct {
		name        string
		payload     map[string]string
		expectedErr error
	}{
		{
			name: "valid",
			payload: map[string]string{
				"Email":    "imhere@mail.com",
				"Password": "secRetec0de",
			},
			expectedErr: nil,
		},
		{
			name: "wrong_email",
			payload: map[string]string{
				"Email":    "wrong_email",
				"Password": "secRetec0de",
			},
			expectedErr: errors.New("incorrect email"),
		},
		{
			name: "wrong_password",
			payload: map[string]string{
				"Email":    "imhere@mail.com",
				"Password": "qwerty123",
			},
			expectedErr: errors.New("incorrect password"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Error(t, tc.expectedErr)
		})
	}
}
