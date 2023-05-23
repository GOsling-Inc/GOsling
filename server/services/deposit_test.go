package services_test

import (
	"testing"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/stretchr/testify/assert"
)

func Test_Deposit_AddDeposit(t *testing.T) {
	db = database.NewMock()
	depos_serv := services.NewDepositService(db)
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
		name        string
		payload     models.Deposit
		expectedErr error
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
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, depos_serv.CreateDeposit(tc.payload))
		})
	}
}

func Test_Deposit_GetUserDeposits(t *testing.T) {
	db = database.NewMock()
	depos_serv := services.NewDepositService(db)
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
		name        string
		payload     string
		deposits    []models.Deposit
		expectedErr error
	}{
		{
			name:        "valid",
			payload:     user0.Id,
			deposits:    []models.Deposit{depos},
			expectedErr: nil,
		},
		{
			name:        "empty_depos",
			payload:     "008",
			deposits:    []models.Deposit(nil),
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d, e := depos_serv.GetUserDeposits(tc.payload)
			assert.Equal(t, tc.deposits, d)
			assert.Equal(t, tc.expectedErr, e)
		})
	}
}
