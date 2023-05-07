package services_test

import (
	"errors"
	"testing"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/stretchr/testify/assert"
)

func Test_User_GetUser(t *testing.T) {
	db = database.NewMock()
	user_serv := services.NewUserService(db)
	db.AddUser(&user0)
	testCases := []struct {
		name        string
		payload     string
		user        models.User
		expectedErr error
	}{
		{
			name:        "valid",
			payload:     user0.Id,
			user:        user0,
			expectedErr: nil,
		},
		{
			name:        "not_found",
			payload:     "otherid",
			user:        models.User{},
			expectedErr: errors.New("not found"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u, e := user_serv.GetUser(tc.payload)
			assert.Equal(t, tc.user, u)
			assert.Equal(t, tc.expectedErr, e)
		})
	}
}

func Test_User_ChangeMainInfo(t *testing.T) {
	db = database.NewMock()
	user_serv := services.NewUserService(db)
	db.AddUser(&user0)
	testCases := []struct {
		name        string
		payload     models.User
		expectedErr error
	}{
		{
			name: "valid",
			payload: models.User{
				Id:        user0.Id,
				Name:      "Name",
				Surname:   "Surname",
				Birthdate: user0.Birthdate,
			},
			expectedErr: nil,
		},
		{
			name: "not_found",
			payload: models.User{
				Id:        "wrongId",
				Name:      "Name",
				Surname:   "Surname",
				Birthdate: user0.Birthdate,
			},
			expectedErr: errors.New("not found"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, user_serv.Change_Main_Info(tc.payload))
		})
	}
}

func Test_User_ChangePassword(t *testing.T) {
	db = database.NewMock()
	user_serv := services.NewUserService(db)
	db.AddUser(&user0)
	testCases := []struct {
		name        string
		payload     models.User
		expectedErr error
	}{
		{
			name: "valid",
			payload: models.User{
				Id:       user0.Id,
				Password: "newpassword",
			},
			expectedErr: nil,
		},
		{
			name: "not_found",
			payload: models.User{
				Id:       "wrongId",
				Password: "newpassword",
			},
			expectedErr: errors.New("not found"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, user_serv.Change_Password(tc.payload))
		})
	}
}
