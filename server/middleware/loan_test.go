package middleware_test

import (
	"testing"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/stretchr/testify/assert"
)

func Test_Loan_ProvideLoan(t *testing.T) {
	db = database.NewMock()
	serv := services.New(db)
	loan_midd := middleware.NewLoanMiddleware(serv)
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
		payload      models.Loan
		expectedCode int
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
			expectedCode: 202,
		},
		{
			name: "unauthorized",
			payload: models.Loan{
				Id:        "007",
				AccountId: "0",
				UserId:    "11",
				Amount:    100,
				Percent:   11,
				State:     "ACTIVE",
			},
			expectedCode: 401,
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
			expectedCode: 500,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := loan_midd.ProvideLoan(tc.payload)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}

func Test_Loan_GetUserLoans(t *testing.T) {
	db = database.NewMock()
	serv := services.New(db)
	loan_midd := middleware.NewLoanMiddleware(serv)
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
		name         string
		payload      string
		loans        []models.Loan
		expectedCode int
	}{
		{
			name:         "valid",
			payload:      user0.Id,
			loans:        []models.Loan{loan},
			expectedCode: 200,
		},
		{
			name:         "empty_depos",
			payload:      "008",
			loans:        []models.Loan(nil),
			expectedCode: 200,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, l, _ := loan_midd.GetUserLoans(tc.payload)
			assert.Equal(t, tc.expectedCode, c)
			assert.Equal(t, tc.loans, l)
		})
	}
}
