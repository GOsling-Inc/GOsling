package services_test

import (
	"testing"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/stretchr/testify/assert"
)

func Test_Insurance_Create_Insurance(t *testing.T) {
	db = database.NewMock()
	ins_serv := services.NewInsuranceService(db)
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
		payload     models.Insurance
		expectedErr error
	}{
		{
			name: "valid",
			payload: models.Insurance{
				Id:        "007",
				AccountId: "0",
				UserId:    user0.Id,
				Amount:    100,
			},
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, ins_serv.CreateInsurance(tc.payload))
		})
	}
}

func Test_Insurance_GetUserInsurances(t *testing.T) {
	db = database.NewMock()
	ins_serv := services.NewInsuranceService(db)
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
		name        string
		payload     string
		insurances  []models.Insurance
		expectedErr error
	}{
		{
			name:        "valid",
			payload:     user0.Id,
			insurances:  []models.Insurance{insurance},
			expectedErr: nil,
		},
		{
			name:        "empty_depos",
			payload:     "008",
			insurances:  []models.Insurance(nil),
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			i, e := ins_serv.GetUserInsurances(tc.payload)
			assert.Equal(t, tc.insurances, i)
			assert.Equal(t, tc.expectedErr, e)
		})
	}
}
