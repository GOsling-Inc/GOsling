package services_test

import (
	"errors"
	"testing"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/stretchr/testify/assert"
)

func Test_Account_AddAccount(t *testing.T) {
	db = database.NewMock()
	acc_serv := services.NewAccountService(db)
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
		name        string
		payload     models.Account
		expectedErr error
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
			expectedErr: nil,
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
			expectedErr: errors.New("can't add an account"),
		},
	}
	t.Run(testCases[0].name, func(t *testing.T) {
		assert.Equal(t, testCases[0].expectedErr, acc_serv.AddAccount(user0.Id, testCases[0].payload))
	})
	for i := 0; i < 11; i++ {
		db.AddAccount(acc)
		acc.Id += "0"
	}
	t.Run(testCases[1].name, func(t *testing.T) {
		assert.Equal(t, testCases[1].expectedErr, acc_serv.AddAccount(user0.Id, testCases[1].payload))
	})
}

func Test_Account_GetAccountbyId(t *testing.T) {
	db = database.NewMock()
	acc_serv := services.NewAccountService(db)
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
			a, e := acc_serv.GetAccountById(tc.payload)
			assert.Equal(t, tc.account, a)
			assert.Equal(t, tc.expectedErr, e)
		})
	}
}

func Test_Account_GetUserAccounts(t *testing.T) {
	db = database.NewMock()
	acc_serv := services.NewAccountService(db)
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
		name        string
		payload     string
		accounts    []models.Account
		expectedErr error
	}{
		{
			name:        "valid",
			payload:     user0.Id,
			accounts:    accs,
			expectedErr: nil,
		},
		{
			name:        "no_accounts",
			payload:     "871",
			accounts:    []models.Account(nil),
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a, e := acc_serv.GetUserAccounts(tc.payload)
			assert.Equal(t, tc.accounts, a)
			assert.Equal(t, tc.expectedErr, e)
		})
	}
}

func Test_Account_DeleteAccount(t *testing.T) {
	db = database.NewMock()
	acc_serv := services.NewAccountService(db)
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
	for i := 0; i < 3; i++ {
		db.AddAccount(acc)
		acc.Id += "0"
	}
	testCases := []struct {
		name        string
		payload     string
		expectedErr error
	}{
		{
			name:        "valid",
			payload:     "00",
			expectedErr: nil,
		},
		{
			name:        "wrong_acc_id",
			payload:     "666",
			expectedErr: errors.New("incorrect account"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, acc_serv.DeleteAccount(tc.payload))
		})
	}
}

func Test_Account_ProvideTransfer(t *testing.T) {
	db = database.NewMock()
	acc_serv := services.NewAccountService(db)
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
		name        string
		payload     models.Trasfer
		expectedErr error
	}{
		{
			name: "not_enough_money",
			payload: models.Trasfer{
				Id:       "007",
				Sender:   acc.Id,
				Receiver: recievers[0].Id,
				Amount:   12345,
			},
			expectedErr: errors.New("sender account error, transaction cancelled"),
		},
		{
			name: "different units",
			payload: models.Trasfer{
				Id:       "007",
				Sender:   acc.Id,
				Receiver: recievers[0].Id,
				Amount:   120,
			},
			expectedErr: errors.New("reciever account error, transaction cancelled"),
		},
		{
			name: "incorrect state",
			payload: models.Trasfer{
				Id:       "007",
				Sender:   acc.Id,
				Receiver: recievers[1].Id,
				Amount:   120,
			},
			expectedErr: errors.New("one of accounts is not active"),
		},
		{
			name: "valid",
			payload: models.Trasfer{
				Id:       "007",
				Sender:   acc.Id,
				Receiver: recievers[2].Id,
				Amount:   120,
			},
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, acc_serv.ProvideTransfer(tc.payload))
		})
	}
}

func Test_Account_ProvideExchange(t *testing.T) {
	db = database.NewMock()
	acc_serv := services.NewAccountService(db)
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
		name        string
		payload     models.Exchange
		expectedErr error
	}{
		{
			name: "not_enough_money",
			payload: models.Exchange{
				Id:             "007",
				Sender:         accs[0].Id,
				Receiver:       reciever.Id,
				SenderAmount:   12345,
				ReceiverAmount: 50,
			},
			expectedErr: errors.New("sender account error, transaction cancelled"),
		},
		{
			name: "same units",
			payload: models.Exchange{
				Id:             "007",
				Sender:         accs[1].Id,
				Receiver:       reciever.Id,
				SenderAmount:   120,
				ReceiverAmount: 50,
			},
			expectedErr: errors.New("reciever account error, transaction cancelled"),
		},
		{
			name: "wrong_user_acc",
			payload: models.Exchange{
				Id:             "007",
				Sender:         accs[2].Id,
				Receiver:       reciever.Id,
				SenderAmount:   120,
				ReceiverAmount: 50,
			},
			expectedErr: errors.New("incorrect account"),
		},
		{
			name: "valid",
			payload: models.Exchange{
				Id:             "007",
				Sender:         accs[0].Id,
				Receiver:       reciever.Id,
				SenderAmount:   120,
				ReceiverAmount: 50,
			},
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, acc_serv.ProvideExchange(tc.payload))
		})
	}
}
