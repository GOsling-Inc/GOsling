package middleware_test

import (
	"testing"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/stretchr/testify/assert"
)

func Test_Insurance_Create_Insurance(t *testing.T) {
	db = database.NewMock()
	ins_serv := services.New(db)
	ins_mid := middleware.NewInsuranceMiddleware(ins_serv)
	db.AddUser(&user0)
	db.AddAccount(models.Account{
		Id:     "0",
		UserId: user0.Id,
		Name:   "Account",
		Type:   "Basic",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	})
	testCases := []struct {
		name         string
		payload      models.Insurance
		expectedCode int
	}{
		{
			name: "valid",
			payload: models.Insurance{
				Id:        "007",
				AccountId: "0",
				UserId:    user0.Id,
				Amount:    100,
			},
			expectedCode: 202,
		},
		{
			name: "wrong_account",
			payload: models.Insurance{
				Id:        "007",
				AccountId: "01",
				UserId:    user0.Id,
				Amount:    100,
			},
			expectedCode: 401,
		},
		{
			name: "wrong_user",
			payload: models.Insurance{
				Id:        "007",
				AccountId: "0",
				UserId:    "666",
				Amount:    100,
			},
			expectedCode: 401,
		},
		{
			name: "not_enough_money",
			payload: models.Insurance{
				Id:        "007",
				AccountId: "0",
				UserId:    user0.Id,
				Part:      1235,
			},
			expectedCode: 401,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := ins_mid.CreateInsurance(tc.payload)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}

func Test_Insurance_GetUserInsurances(t *testing.T) {
	db = database.NewMock()
	ins_serv := services.New(db)
	ins_mid := middleware.NewInsuranceMiddleware(ins_serv)
	db.AddUser(&user0)
	db.AddAccount(models.Account{
		Id:     "0",
		UserId: user0.Id,
		Name:   "Account",
		Type:   "Basic",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	})
	insurance := models.Insurance{
		Id:        "007",
		AccountId: "0",
		UserId:    user0.Id,
		Amount:    100,
	}
	db.AddInsurance(insurance)
	testCases := []struct {
		name         string
		payload      string
		insurances   []models.Insurance
		expectedCode int
	}{
		{
			name:         "valid",
			payload:      user0.Id,
			insurances:   []models.Insurance{insurance},
			expectedCode: 200,
		},
		{
			name:         "empty_depos",
			payload:      "008",
			insurances:   []models.Insurance(nil),
			expectedCode: 200,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, i, _ := ins_mid.GetUserInsurances(tc.payload)
			assert.Equal(t, tc.expectedCode, c)
			assert.Equal(t, tc.insurances, i)
		})
	}
}
