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

func Test_Account_AddAccount(t *testing.T) {
	db = database.NewMock()
	acc_serv := services.New(db)
	acc_mid := middleware.NewAccountMiddleware(acc_serv)
	acc := models.Account{
		Id:     "0",
		UserId: user0.Id,
		Name:   "First_Account",
		Type:   "Basic",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	db.AddUser(&user0)
	testCases := []struct {
		name         string
		payload      models.Account
		expectedCode int
	}{
		{
			name: "valid",
			payload: models.Account{
				Id:     "1",
				UserId: user0.Id,
				Name:   "First_Account",
				Type:   "Basic",
				Unit:   "BYN",
				Amount: 123,
				State:  "Active",
			},
			expectedCode: 201,
		},
		{
			name: "too_much_accs",
			payload: models.Account{
				Id:     "666",
				UserId: user0.Id,
				Name:   "Beta_Account",
				Type:   "Basic",
				Unit:   "BYN",
				Amount: 123,
				State:  "Active",
			},
			expectedCode: 500,
		},
	}
	t.Run(testCases[0].name, func(t *testing.T) {
		c, _ := acc_mid.AddAccount(user0.Id, testCases[0].payload)
		assert.Equal(t, testCases[0].expectedCode, c)
	})
	for i := 0; i < 11; i++ {
		db.AddAccount(acc)
		acc.Id += "0"
	}
	t.Run(testCases[1].name, func(t *testing.T) {
		c, _ := acc_mid.AddAccount(user0.Id, testCases[1].payload)
		assert.Equal(t, testCases[1].expectedCode, c)
	})
}

func Test_Account_GetAccountbyId(t *testing.T) {
	db = database.NewMock()
	acc_serv := services.New(db)
	acc_mid := middleware.NewAccountMiddleware(acc_serv)
	acc := models.Account{
		Id:     "0",
		UserId: user0.Id,
		Name:   "First_Account",
		Type:   "Basic",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	db.AddAccount(acc)
	testCases := []struct {
		name        string
		payload     string
		account     models.Account
		expectedErr error
	}{
		{
			name:        "valid",
			payload:     "0",
			account:     acc,
			expectedErr: nil,
		},
		{
			name:        "wrong_id",
			payload:     "iamwrong",
			account:     models.Account{},
			expectedErr: errors.New("not found"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a, e := acc_mid.GetAccountById(tc.payload)
			assert.Equal(t, tc.account, a)
			assert.Equal(t, tc.expectedErr, e)
		})
	}
}

func Test_Account_GetUserAccounts(t *testing.T) {
	db = database.NewMock()
	acc_serv := services.New(db)
	acc_mid := middleware.NewAccountMiddleware(acc_serv)
	acc := models.Account{
		Id:     "0",
		UserId: user0.Id,
		Name:   "First_Account",
		Type:   "Basic",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	db.AddUser(&user0)
	accs := []models.Account{}
	for i := 0; i < 5; i++ {
		db.AddAccount(acc)
		accs = append(accs, acc)
		acc.Id += "0"
	}
	testCases := []struct {
		name         string
		payload      string
		accounts     []models.Account
		expectedCode int
	}{
		{
			name:         "valid",
			payload:      user0.Id,
			accounts:     accs,
			expectedCode: 200,
		},
		{
			name:         "no_accounts",
			payload:      "871",
			accounts:     []models.Account(nil),
			expectedCode: 200,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, a, _ := acc_mid.GetUserAccounts(tc.payload)
			assert.Equal(t, tc.accounts, a)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}

func Test_Account_DeleteAccount(t *testing.T) {
	db = database.NewMock()
	acc_serv := services.New(db)
	acc_mid := middleware.NewAccountMiddleware(acc_serv)
	acc := models.Account{
		Id:     "0",
		UserId: user0.Id,
		Name:   "First_Account",
		Type:   "Basic",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	user0.Password, _ = acc_serv.Hash(user0.Password)
	db.AddUser(&user0)
	for i := 0; i < 3; i++ {
		db.AddAccount(acc)
		acc.Id += "0"
	}
	testCases := []struct {
		name         string
		password     string
		user         string
		acc          string
		expectedCode int
	}{
		{
			name:         "wrong_user_id",
			password:     "secRetec0de",
			user:         "666",
			acc:          "666",
			expectedCode: 401,
		},
		{
			name:         "wrong_password",
			password:     "password",
			user:         "666",
			acc:          "666",
			expectedCode: 401,
		},
		{
			name:         "wrong_acc_id",
			password:     "secRetec0de",
			user:         user0.Id,
			acc:          "666",
			expectedCode: 401,
		},
		{
			name:         "valid",
			password:     "secRetec0de",
			user:         user0.Id,
			acc:          "00",
			expectedCode: 202,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := acc_mid.DeleteAccount(tc.user, tc.acc, tc.password)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}

func Test_Account_ProvideTransfer(t *testing.T) {
	db = database.NewMock()
	acc_serv := services.New(db)
	acc_mid := middleware.NewAccountMiddleware(acc_serv)
	acc := models.Account{
		Id:     "0",
		UserId: user0.Id,
		Name:   "First_Account",
		Type:   "Basic",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	db.AddUser(&user0)
	db.AddAccount(acc)
	recievers := []models.Account{
		{
			Id:     "1",
			UserId: "787",
			Unit:   "USD",
			Amount: 12,
			State:  "ACTIVE",
		},
		{
			Id:     "2",
			UserId: "787",
			Unit:   "BYN",
			Amount: 12,
			State:  "FROZEN",
		},
		{
			Id:     "3",
			UserId: "787",
			Unit:   "BYN",
			Amount: 12,
			State:  "ACTIVE",
		},
	}
	for _, r := range recievers {
		db.AddAccount(r)
	}
	testCases := []struct {
		name         string
		id           string
		payload      models.Trasfer
		expectedCode int
	}{
		{
			name: "incorrect_acc",
			id:   "666",
			payload: models.Trasfer{
				Id:       "007",
				Sender:   acc.Id,
				Receiver: recievers[2].Id,
				Amount:   120,
			},
			expectedCode: 401,
		},
		{
			name: "not_enough_money",
			id:   user0.Id,
			payload: models.Trasfer{
				Id:       "007",
				Sender:   acc.Id,
				Receiver: recievers[0].Id,
				Amount:   12345,
			},
			expectedCode: 500,
		},
		{
			name: "different units",
			id:   user0.Id,
			payload: models.Trasfer{
				Id:       "007",
				Sender:   acc.Id,
				Receiver: recievers[0].Id,
				Amount:   120,
			},
			expectedCode: 500,
		},
		{
			name: "incorrect state",
			id:   user0.Id,
			payload: models.Trasfer{
				Id:       "007",
				Sender:   acc.Id,
				Receiver: recievers[1].Id,
				Amount:   120,
			},
			expectedCode: 500,
		},
		{
			name: "valid",
			id:   user0.Id,
			payload: models.Trasfer{
				Id:       "007",
				Sender:   acc.Id,
				Receiver: recievers[2].Id,
				Amount:   120,
			},
			expectedCode: 202,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := acc_mid.ProvideTransfer(tc.id, tc.payload)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}

func Test_Account_ProvideExchange(t *testing.T) {
	db = database.NewMock()
	acc_serv := services.New(db)
	acc_mid := middleware.NewAccountMiddleware(acc_serv)
	reciever := models.Account{
		Id:     "0",
		UserId: user0.Id,
		Unit:   "USD",
		Amount: 123,
		State:  "ACTIVE",
	}
	db.AddAccount(reciever)
	db.AddUser(&user0)
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
			UserId: "wrong_id",
			Unit:   "BYN",
			Amount: 1234,
			State:  "ACTIVE",
		},
	}
	for _, r := range accs {
		db.AddAccount(r)
	}
	testCases := []struct {
		name         string
		id           string
		payload      models.Exchange
		expectedCode int
	}{
		{
			name: "not_enough_money",
			id:   user0.Id,
			payload: models.Exchange{
				Id:             "007",
				Sender:         accs[0].Id,
				Receiver:       reciever.Id,
				SenderAmount:   12345,
				ReceiverAmount: 50,
			},
			expectedCode: 500,
		},
		{
			name: "same units",
			id:   user0.Id,
			payload: models.Exchange{
				Id:             "007",
				Sender:         accs[1].Id,
				Receiver:       reciever.Id,
				SenderAmount:   120,
				ReceiverAmount: 50,
			},
			expectedCode: 500,
		},
		{
			name: "wrong_user_acc",
			id:   user0.Id,
			payload: models.Exchange{
				Id:             "007",
				Sender:         accs[2].Id,
				Receiver:       reciever.Id,
				SenderAmount:   120,
				ReceiverAmount: 50,
			},
			expectedCode: 401,
		},
		{
			name: "valid",
			id:   user0.Id,
			payload: models.Exchange{
				Id:             "007",
				Sender:         accs[0].Id,
				Receiver:       reciever.Id,
				SenderAmount:   120,
				ReceiverAmount: 50,
			},
			expectedCode: 200,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := acc_mid.ProvideExchange(tc.id, tc.payload)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}
