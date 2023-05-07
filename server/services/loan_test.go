package services_test

import (
	"errors"
	"testing"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/stretchr/testify/assert"
)

func Test_Loan_ProvideLoan(t *testing.T) {
	db = database.NewMock()
	loan_serv := services.NewLoanService(db)
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
		payload     models.Loan
		expectedErr error
	}{
		{
			name: "valid",
			payload: models.Loan{
				Id:        "007",
				AccountId: "0",
				UserId:    user0.Id,
				Amount:    100,
				Percent:   11,
				State:     "ACTIVE",
			},
			expectedErr: nil,
		},
		{
			name: "has_opened_loans",
			payload: models.Loan{
				Id:        "007",
				AccountId: "0",
				UserId:    user0.Id,
				Amount:    100,
				Percent:   11,
				State:     "ACTIVE",
			},
			expectedErr: errors.New("can't open a new loan"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, loan_serv.ProvideLoan(tc.payload))
		})
	}
}

func Test_Loan_GetUserLoans(t *testing.T) {
	db = database.NewMock()
	loan_serv := services.NewLoanService(db)
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
	loan := models.Loan{
		Id:        "007",
		AccountId: "0",
		UserId:    user0.Id,
		Amount:    100,
		Percent:   11,
	}
	db.AddLoan(loan)
	testCases := []struct {
		name        string
		payload     string
		loans       []models.Loan
		expectedErr error
	}{
		{
			name:        "valid",
			payload:     user0.Id,
			loans:       []models.Loan{loan},
			expectedErr: nil,
		},
		{
			name:        "empty_depos",
			payload:     "008",
			loans:       []models.Loan(nil),
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l, e := loan_serv.GetUserLoans(tc.payload)
			assert.Equal(t, tc.loans, l)
			assert.Equal(t, tc.expectedErr, e)
		})
	}
}
