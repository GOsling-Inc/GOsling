package middleware_test

import (
	"testing"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/stretchr/testify/assert"
)

func Test_Manage_GetConfirms(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.New(db)
	manage_mid := middleware.NewManagerMiddleware(manage_serv)
	db.AddUser(&user0)
	acc := models.Account{
		Id:     "0",
		UserId: user0.Id,
		Name:   "Account",
		Type:   "BASIC",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	db.AddAccount(acc)
	db.AddLoan(models.Loan{
		Id:        "007",
		AccountId: acc.Id,
		UserId:    user0.Id,
		Amount:    100,
		Percent:   11,
		State:     "PENDING",
	})
	db.AddInsurance(models.Insurance{
		Id:        "007",
		AccountId: acc.Id,
		UserId:    user0.Id,
		Amount:    100,
		State:     "PENDING",
	})
	db.AddDeposit(models.Deposit{
		Id:        "007",
		AccountId: acc.Id,
		UserId:    user0.Id,
		Amount:    100,
		Percent:   11,
		State:     "PENDING",
	})
	res := []models.Unconfirmed{
		{
			Table: "loans",
			Id:    "007",
		},
		{
			Table: "deposits",
			Id:    "007",
		},
		{
			Table: "insurances",
			Id:    "007",
		},
	}
	testCases := []struct {
		name        string
		expectedRes []models.Unconfirmed
	}{
		{
			name:        "valid",
			expectedRes: res,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			code, c := manage_mid.GetConfirms()
			assert.Equal(t, tc.expectedRes, c)
			assert.Equal(t, 200, code)
		})
	}
}

func Test_Manage_Confirm(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.New(db)
	manage_mid := middleware.NewManagerMiddleware(manage_serv)
	db.AddUser(&user0)
	acc := models.Account{
		Id:     "0",
		UserId: user0.Id,
		Name:   "Account",
		Type:   "BASIC",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	db.AddAccount(acc)
	db.AddLoan(models.Loan{
		Id:        "007",
		AccountId: acc.Id,
		UserId:    user0.Id,
		Amount:    100,
		Percent:   11,
		State:     "PENDING",
	})
	db.AddDeposit(models.Deposit{
		Id:        "007",
		AccountId: acc.Id,
		UserId:    user0.Id,
		Amount:    100,
		Percent:   11,
		State:     "PENDING",
	})
	db.AddInsurance(models.Insurance{
		Id:        "007",
		AccountId: acc.Id,
		UserId:    user0.Id,
		Amount:    100,
		State:     "PENDING",
	})
	testCases := []struct {
		name         string
		id           string
		table        string
		state        string
		expectedCode int
	}{
		{
			name:         "loan_valid_confirm",
			table:        "loans",
			id:           "007",
			state:        "CONFIRMED",
			expectedCode: 200,
		},
		{
			name:         "loan_valid_denied",
			table:        "loans",
			id:           "007",
			state:        "DENIED",
			expectedCode: 200,
		},
		{
			name:         "wrong_loan_id",
			table:        "loans",
			id:           "666",
			state:        "CONFIRMED",
			expectedCode: 500,
		},
		{
			name:         "deposit_valid_confirm",
			table:        "deposits",
			id:           "007",
			state:        "CONFIRMED",
			expectedCode: 200,
		},
		{
			name:         "deposit_valid_denied",
			table:        "deposits",
			id:           "007",
			state:        "DENIED",
			expectedCode: 200,
		},
		{
			name:         "wrong_deposit_id",
			table:        "deposits",
			id:           "666",
			state:        "CONFIRMED",
			expectedCode: 500,
		},
		{
			name:         "insurance_valid_confirm",
			table:        "insurances",
			id:           "007",
			state:        "CONFIRMED",
			expectedCode: 200,
		},
		{
			name:         "insurances_valid_denied",
			table:        "insurances",
			id:           "007",
			state:        "DENIED",
			expectedCode: 200,
		},
		{
			name:         "wrong_insurance_id",
			table:        "insurances",
			id:           "666",
			state:        "CONFIRMED",
			expectedCode: 500,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := manage_mid.Confirm(tc.id, tc.table, tc.state)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}

func Test_Manage_GetAccounts(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.New(db)
	manage_mid := middleware.NewManagerMiddleware(manage_serv)
	db.AddUser(&user0)
	res := make([]models.Account, 10)
	accs := []models.Account{
		{
			Id:     "1",
			UserId: user0.Id,
			Unit:   "BYN",
			Amount: 1234,
			State:  "ACTIVE",
		},
		{
			Id:     "2",
			UserId: user0.Id,
			Unit:   "USD",
			Amount: 1234,
			State:  "ACTIVE",
		},
		{
			Id:     "3",
			UserId: user0.Id,
			Unit:   "EUR",
			Amount: 1234,
			State:  "ACTIVE",
		},
	}
	for _, r := range accs {
		db.AddAccount(r)
		res = append(res, r)
	}
	testCases := []struct {
		name         string
		expectedAccs []models.Account
	}{
		{
			name:         "valid",
			expectedAccs: res,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, a := manage_mid.GetAccounts()
			assert.Equal(t, c, 200)
			assert.Equal(t, tc.expectedAccs, a)
		})
	}
}

func Test_Manage_UpdateAccount(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.New(db)
	manage_mid := middleware.NewManagerMiddleware(manage_serv)
	db.AddUser(&user0)
	acc := models.Account{
		Id:     "0",
		UserId: user0.Id,
		Name:   "Account",
		Type:   "BASIC",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	db.AddAccount(acc)
	testCases := []struct {
		name         string
		id           string
		state        string
		expectedCode int
	}{
		{
			name:         "undefined_state",
			id:           "0",
			state:        "ABOBA",
			expectedCode: 500,
		},
		{
			name:         "not_fount",
			id:           "777",
			state:        "FROZEN",
			expectedCode: 500,
		},
		{
			name:         "valid",
			id:           "0",
			state:        "FREEZED",
			expectedCode: 200,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := manage_mid.UpdateAccount(tc.id, tc.state)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}

func Test_Manage_GetTransactions(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.New(db)
	manage_mid := middleware.NewManagerMiddleware(manage_serv)
	db.AddUser(&user0)
	acc := models.Account{
		Id:     "0",
		UserId: user0.Id,
		Name:   "Account",
		Type:   "BASIC",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	db.AddAccount(acc)
	res := make([]models.Trasfer, 10)
	trans := []models.Trasfer{
		{
			Id:       "1",
			Sender:   acc.Id,
			Receiver: "123",
			Amount:   1234,
		},
		{
			Id:       "2",
			Sender:   acc.Id,
			Receiver: "123",
			Amount:   1,
		},
		{
			Id:       "3",
			Sender:   acc.Id,
			Receiver: "123",
			Amount:   123,
		},
	}
	for _, r := range trans {
		db.AddTransfer(r)
		res = append(res, r)
	}
	testCases := []struct {
		name          string
		expectedTrans []models.Trasfer
	}{
		{
			name:          "valid",
			expectedTrans: res,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, tr := manage_mid.GetTransactions()
			assert.Equal(t, c, 200)
			assert.Equal(t, tc.expectedTrans, tr)
		})
	}
}
func Test_Manage_CancelTransaction(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.New(db)
	manage_mid := middleware.NewManagerMiddleware(manage_serv)
	db.AddUser(&user0)
	acc := models.Account{
		Id:     "0",
		UserId: user0.Id,
		Name:   "Account",
		Type:   "BASIC",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	reciever := models.Account{

		Id:     "1",
		UserId: "787",
		Unit:   "BYN",
		Amount: 12,
		State:  "ACTIVE",
	}
	db.AddAccount(acc)
	db.AddAccount(reciever)
	db.AddTransfer(models.Trasfer{
		Id:       "007",
		Sender:   acc.Id,
		Receiver: reciever.Id,
		Amount:   120,
	})
	testCases := []struct {
		name         string
		payload      string
		expectedCode int
	}{
		{
			name:         "wrong_id",
			payload:      "666",
			expectedCode: 500,
		},
		{
			name:         "valid",
			payload:      "007",
			expectedCode: 200,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := manage_mid.CancelTransaction(tc.payload)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}

func Test_Manage_GetUsers(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.New(db)
	manage_mid := middleware.NewManagerMiddleware(manage_serv)
	res := make([]models.User, 10)
	users := []models.User{
		{
			Id:        "1",
			Name:      "ABOBA",
			Surname:   "ABOBOV",
			Email:     "test1@gmail.com",
			Password:  "154492Ad",
			Role:      "user",
			Birthdate: "2002-01-01",
		},
		{
			Id:        "2",
			Name:      "Pavel",
			Surname:   "Petrov",
			Email:     "test2@gmail.com",
			Password:  "15817GrAb",
			Role:      "user",
			Birthdate: "2006-08-11",
		},
		{
			Id:        "3",
			Name:      "USER",
			Surname:   "NEW",
			Email:     "test3@gmail.com",
			Password:  "11111111",
			Role:      "user",
			Birthdate: "2012-11-01",
		},
	}
	for _, u := range users {
		db.AddUser(&u)
		res = append(res, u)
	}
	testCases := []struct {
		name          string
		ExpectedUsers []models.User
	}{
		{
			name:          "valid",
			ExpectedUsers: res,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, u := manage_mid.GetUsers()
			assert.Equal(t, c, 200)
			assert.Equal(t, tc.ExpectedUsers, u)
		})
	}
}

func Test_Manage_UpdateRole(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.New(db)
	manage_mid := middleware.NewManagerMiddleware(manage_serv)
	db.AddUser(&user0)
	testCases := []struct {
		name         string
		id           string
		role         string
		expectedCode int
	}{
		{
			name:         "valid",
			id:           user0.Id,
			role:         "manager",
			expectedCode: 200,
		},
		{
			name:         "wrong_id",
			id:           "wrong_id",
			role:         "user",
			expectedCode: 500,
		},
		{
			name:         "wrong_role",
			id:           user0.Id,
			role:         "fidget_spiner",
			expectedCode: 500,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := manage_mid.UpdateRole(tc.id, tc.role)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}
