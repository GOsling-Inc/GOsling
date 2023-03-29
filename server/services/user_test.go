package services_test

import (
	"testing"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"

	"github.com/stretchr/testify/assert"
)

var (
	serv services.UserService
)

func Test_User_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *models.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *models.User {
				return &models.User{
					Email:    "examle@gmail.com",
					Password: "12345678",
				}
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *models.User {
				return &models.User{
					Email:    "",
					Password: "12345678",
				}
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *models.User {
				return &models.User{
					Email:    "email",
					Password: "12345678",
				}
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *models.User {
				return &models.User{
					Email:    "examle@gmail.com",
					Password: "",
				}
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *models.User {
				return &models.User{
					Email:    "examle@gmail.com",
					Password: "123",
				}
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, serv.Validate(tc.u()))
			} else {
				assert.Error(t, serv.Validate(tc.u()))
			}
		})
	}
}
