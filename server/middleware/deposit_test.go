package middleware_test

import (
	"testing"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/stretchr/testify/assert"
)

func Test_Deposit_CreateDeposit(t *testing.T) {
	db = database.NewMock()
	depos_serv := services.New(db)
	depos_midd := middleware.NewDepositMiddleware(depos_serv)
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
		payload      models.Deposit
		expectedCode int
	}{
		{
			name: "valid",
			payload: models.Deposit{
				Id:        "007",
				AccountId: "0",
				UserId:    user0.Id,
				Amount:    100,
				Percent:   11,
			},
			expectedCode: 202,
		},
		{
			name: "wrong_account",
			payload: models.Deposit{
				Id:        "007",
				AccountId: "01",
				UserId:    user0.Id,
				Amount:    100,
				Percent:   11,
			},
			expectedCode: 401,
		},
		{
			name: "wrong_user",
			payload: models.Deposit{
				Id:        "007",
				AccountId: "0",
				UserId:    "666",
				Amount:    100,
				Percent:   11,
			},
			expectedCode: 401,
		},
		{
			name: "not_enough_money",
			payload: models.Deposit{
				Id:        "007",
				AccountId: "0",
				UserId:    user0.Id,
				Amount:    100000,
				Percent:   11,
			},
			expectedCode: 401,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := depos_midd.CreateDeposit(tc.payload)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}

func Test_Deposit_GetUserDeposits(t *testing.T) {
	db = database.NewMock()
	depos_serv := services.New(db)
	depos_mid := middleware.NewDepositMiddleware(depos_serv)
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
	depos := models.Deposit{
		Id:        "007",
		AccountId: "0",
		UserId:    user0.Id,
		Amount:    100,
		Percent:   11,
	}
	db.AddDeposit(depos)
	testCases := []struct {
		name         string
		payload      string
		deposits     []models.Deposit
		expectedCode int
	}{
		{
			name:         "valid",
			payload:      user0.Id,
			deposits:     []models.Deposit{depos},
			expectedCode: 200,
		},
		{
			name:         "empty_depos",
			payload:      "008",
			deposits:     []models.Deposit(nil),
			expectedCode: 200,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, d, _ := depos_mid.GetUserDeposits(tc.payload)
			assert.Equal(t, tc.expectedCode, c)
			assert.Equal(t, tc.deposits, d)
		})
	}
}
