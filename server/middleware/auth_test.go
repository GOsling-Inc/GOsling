package middleware_test

import (
	"errors"
	"testing"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"

	"github.com/stretchr/testify/assert"
)

var (
	db    *database.MockDatabase
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

func Test_Auth_Validate(t *testing.T) {
	db = database.NewMock()
	serv := services.New(db)
	auth_midd := middleware.NewAuthMiddleware(serv)
	testCases := []struct {
		name    string
		u       models.User
		isValid bool
	}{
		{
			name: "valid",
			u: models.User{
				Email:    "imhere@mail.com",
				Password: "secRetec0de",
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: models.User{
				Email:    "",
				Password: "12345678",
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: models.User{
				Email:    "email",
				Password: "12345678",
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: models.User{
				Email:    "examle@gmail.com",
				Password: "",
			},
			isValid: false,
		},
		{
			name: "short password",
			u: models.User{
				Email:    "examle@gmail.com",
				Password: "123",
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, auth_midd.Validate(tc.u))
			} else {
				assert.Error(t, auth_midd.Validate(tc.u))
			}
		})
	}
}

func Test_Auth_SighIn(t *testing.T) {
	db = database.NewMock()
	serv := services.New(db)
	auth_midd := middleware.NewAuthMiddleware(serv)
	user := models.User{
		Id:        "1234",
		Name:      "User",
		Surname:   "Existing",
		Email:     "imhere@mail.com",
		Password:  "secRetec0de",
		Role:      "user",
		Birthdate: "1979-09-21",
	}
	user.Password, _ = serv.Hash(user.Password)
	db.AddUser(&user)
	testCases := []struct {
		name         string
		payload      models.User
		expectedCode int
	}{
		{
			name: "valid",
			payload: models.User{
				Email:    "imhere@mail.com",
				Password: "secRetec0de",
			},
			expectedCode: 200,
		},
		{
			name: "empty_field",
			payload: models.User{
				Email:    "imhere@mail.com",
				Password: "",
			},
			expectedCode: 401,
		},
		{
			name: "wrong_email",
			payload: models.User{
				Email:    "wrong_email@aboba.com",
				Password: "secRetec0de",
			},
			expectedCode: 401,
		},
		{
			name: "wrong_password",
			payload: models.User{
				Email:    "imhere@mail.com",
				Password: "wr0nG_pa$$w0rD",
			},
			expectedCode: 401,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := auth_midd.SignIn(&tc.payload)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}

func Test_Auth_SignUp(t *testing.T) {
	db = database.NewMock()
	serv := services.New(db)
	auth_midd := middleware.NewAuthMiddleware(serv)
	user0.Password, _ = serv.Hash(user0.Password)
	db.AddUser(&user0)
	testCases := []struct {
		name         string
		payload      models.User
		expectedCode int
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
			expectedCode: 201,
		},
		{
			name: "wrong_info",
			payload: models.User{
				Id:        "",
				Name:      "Same",
				Surname:   "Email",
				Email:     "imhere",
				Password:  "1544",
				Role:      "user",
				Birthdate: "2002-01-01",
			},
			expectedCode: 401,
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
			expectedCode: 401,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := auth_midd.SignUp(&tc.payload)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}

func Test_Auth_AuthManager(t *testing.T) {
	db = database.NewMock()
	serv := services.New(db)
	auth_midd := middleware.NewAuthMiddleware(serv)
	user0.Password, _ = serv.Hash(user0.Password)

	db.AddUser(&user0)
	db.AddUser(&models.User{
		Id:        "007",
		Name:      "ABOBA",
		Surname:   "ABOBOV",
		Email:     "test1@gmail.com",
		Password:  "154492Ad",
		Role:      "MANAGER",
		Birthdate: "2002-01-01",
	})
	testCases := []struct {
		name        string
		payload     string
		expectedErr error
	}{
		{
			name:        "valid",
			payload:     "007",
			expectedErr: nil,
		},
		{
			name:        "invalid_id",
			payload:     "777",
			expectedErr: errors.New("incorrect id"),
		},
		{
			name:        "invalid_role",
			payload:     user0.Id,
			expectedErr: errors.New("access denied"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, auth_midd.AuthManager(tc.payload))
		})
	}
}

func Test_Auth_AuthAdmin(t *testing.T) {
	db = database.NewMock()
	serv := services.New(db)
	auth_midd := middleware.NewAuthMiddleware(serv)
	user0.Password, _ = serv.Hash(user0.Password)

	db.AddUser(&user0)
	db.AddUser(&models.User{
		Id:        "007",
		Name:      "ABOBA",
		Surname:   "ABOBOV",
		Email:     "test1@gmail.com",
		Password:  "154492Ad",
		Role:      "MANAGER",
		Birthdate: "2002-01-01",
	})
	db.AddUser(&models.User{
		Id:        "700",
		Name:      "ABIBA",
		Surname:   "ABIBOV",
		Email:     "test1@gmail.com",
		Password:  "154492Ad",
		Role:      "admin",
		Birthdate: "2002-01-01",
	})
	testCases := []struct {
		name        string
		payload     string
		expectedErr error
	}{
		{
			name:        "valid",
			payload:     "700",
			expectedErr: nil,
		},
		{
			name:        "invalid_id",
			payload:     "777",
			expectedErr: errors.New("incorrect id"),
		},
		{
			name:        "invalid_role1",
			payload:     user0.Id,
			expectedErr: errors.New("access denied"),
		},
		{
			name:        "invalid_role2",
			payload:     "007",
			expectedErr: errors.New("access denied"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, auth_midd.AuthAdmin(tc.payload))
		})
	}
}
