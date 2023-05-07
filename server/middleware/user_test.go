package middleware_test

import (
	"testing"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/stretchr/testify/assert"
)

func Test_User_GetUser(t *testing.T) {
	db = database.NewMock()
	serv := services.New(db)
	user_midd := middleware.NewUserMiddleware(serv)
	db.AddUser(&user0)
	testCases := []struct {
		name         string
		payload      string
		user         models.User
		expectedCode int
	}{
		{
			name:         "valid",
			payload:      user0.Id,
			user:         user0,
			expectedCode: 200,
		},
		{
			name:         "not_found",
			payload:      "otherid",
			user:         models.User{},
			expectedCode: 401,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, u, _ := user_midd.GetUser(tc.payload)
			assert.Equal(t, tc.user, u)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}

func Test_User_ChangeMainInfo(t *testing.T) {
	db = database.NewMock()
	serv := services.New(db)
	user_midd := middleware.NewUserMiddleware(serv)
	db.AddUser(&user0)
	testCases := []struct {
		name         string
		payload      models.User
		expectedCode int
	}{
		{
			name: "valid",
			payload: models.User{
				Id:        user0.Id,
				Name:      "Name",
				Surname:   "Surname",
				Birthdate: user0.Birthdate,
			},
			expectedCode: 202,
		},
		{
			name: "not_found",
			payload: models.User{
				Id:        "wrongId",
				Name:      "Name",
				Surname:   "Surname",
				Birthdate: user0.Birthdate,
			},
			expectedCode: 401,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := user_midd.Change_Main_Info(tc.payload)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}

func Test_User_ChangePassword(t *testing.T) {
	db = database.NewMock()
	serv := services.New(db)
	user_midd := middleware.NewUserMiddleware(serv)
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
		userid       string
		oldPassword  string
		newPassword  string
		expectedCode int
	}{
		{
			name:         "valid",
			userid:       user.Id,
			oldPassword:  "secRetec0de",
			newPassword:  "newPa$$w0rd",
			expectedCode: 202,
		},
		{
			name:         "not_found",
			userid:       "666",
			oldPassword:  user0.Password,
			newPassword:  "newPa$$w0rd",
			expectedCode: 401,
		},
		{
			name:         "password_not_match",
			userid:       user0.Id,
			oldPassword:  "ABOBABOBABOBA",
			newPassword:  "newPa$$w0rd",
			expectedCode: 401,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := user_midd.Change_Password(tc.userid, tc.oldPassword, tc.newPassword)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}
