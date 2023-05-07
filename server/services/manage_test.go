package services_test

import (
	"errors"
	"testing"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/stretchr/testify/assert"
)

func Test_Manage_GetConfirms(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.NewManageService(db)
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
			assert.Equal(t, tc.expectedRes, manage_serv.GetConfirms())
		})
	}
}

func Test_Manage_ConfirmLoan(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.NewManageService(db)
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
	testCases := []struct {
		name        string
		id          string
		state       string
		expectedErr error
	}{
		{
			name:        "valid_confirm",
			id:          "007",
			state:       "CONFIRMED",
			expectedErr: nil,
		},
		{
			name:        "valid_denied",
			id:          "007",
			state:       "DENIED",
			expectedErr: nil,
		},
		{
			name:        "wrong_loan_id",
			id:          "666",
			state:       "CONFIRMED",
			expectedErr: errors.New("not found"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, manage_serv.ConfirmLoan(tc.id, tc.state))
		})
	}
}

func Test_Manage_ConfirmDeposit(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.NewManageService(db)
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
	db.AddDeposit(models.Deposit{
		Id:        "007",
		AccountId: acc.Id,
		UserId:    user0.Id,
		Amount:    100,
		Percent:   11,
		State:     "PENDING",
	})
	testCases := []struct {
		name        string
		id          string
		state       string
		expectedErr error
	}{
		{
			name:        "valid_confirm",
			id:          "007",
			state:       "CONFIRMED",
			expectedErr: nil,
		},
		{
			name:        "valid_denied",
			id:          "007",
			state:       "DENIED",
			expectedErr: nil,
		},
		{
			name:        "wrong_deposit_id",
			id:          "666",
			state:       "CONFIRMED",
			expectedErr: errors.New("not found"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, manage_serv.ConfirmDeposit(tc.id, tc.state))
		})
	}
}

func Test_Manage_ConfirmInsurance(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.NewManageService(db)
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
	db.AddInsurance(models.Insurance{
		Id:        "007",
		AccountId: acc.Id,
		UserId:    user0.Id,
		Amount:    100,
		State:     "PENDING",
	})
	testCases := []struct {
		name        string
		id          string
		state       string
		expectedErr error
	}{
		{
			name:        "valid_confirm",
			id:          "007",
			state:       "CONFIRMED",
			expectedErr: nil,
		},
		{
			name:        "valid_denied",
			id:          "007",
			state:       "DENIED",
			expectedErr: nil,
		},
		{
			name:        "wrong_insurance_id",
			id:          "666",
			state:       "CONFIRMED",
			expectedErr: errors.New("not found"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, manage_serv.ConfirmInsurance(tc.id, tc.state))
		})
	}
}

func Test_Manage_GetAccounts(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.NewManageService(db)
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
			assert.Equal(t, tc.expectedAccs, manage_serv.GetAccounts())
		})
	}
}

func Test_Manage_UpdateAccount(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.NewManageService(db)
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
		name        string
		id          string
		state       string
		expectedErr error
	}{
		{
			name:        "valid",
			id:          "0",
			state:       "FROZEN",
			expectedErr: nil,
		},
		{
			name:        "not_fount",
			id:          "777",
			state:       "FROZEN",
			expectedErr: errors.New("not found"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, manage_serv.UpdateAccount(tc.id, tc.state))
		})
	}
}

func Test_Manage_GetTransactions(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.NewManageService(db)
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
			assert.Equal(t, tc.expectedTrans, manage_serv.GetTransactions())
		})
	}
}

func Test_Manage_GetTransferbyId(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.NewManageService(db)
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
	}
	testCases := []struct {
		name        string
		id          string
		transfer    models.Trasfer
		expectedErr error
	}{
		{
			name:        "valid",
			id:          "2",
			transfer:    trans[1],
			expectedErr: nil,
		},
		{
			name:        "not_found",
			id:          "777",
			transfer:    models.Trasfer{},
			expectedErr: errors.New("not found"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tr, e := manage_serv.GetTransferById(tc.id)
			assert.Equal(t, tc.expectedErr, e)
			assert.Equal(t, tc.transfer, tr)
		})
	}
}

func Test_Manage_CancelTransaction(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.NewManageService(db)
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
	testCases := []struct {
		name        string
		payload     models.Trasfer
		expectedErr error
	}{
		{
			name: "valid",
			payload: models.Trasfer{
				Id:       "007",
				Sender:   acc.Id,
				Receiver: reciever.Id,
				Amount:   120,
			},
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, manage_serv.CancelTransaction(tc.payload))
		})
	}
}

func Test_Manage_GetUsers(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.NewManageService(db)
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
			assert.Equal(t, tc.ExpectedUsers, manage_serv.GetUsers())
		})
	}
}

func Test_Manage_UpdateRole(t *testing.T) {
	db = database.NewMock()
	manage_serv := services.NewManageService(db)
	db.AddUser(&user0)
	testCases := []struct {
		name        string
		id          string
		role        string
		expectedErr error
	}{
		{
			name:        "valid",
			id:          user0.Id,
			role:        "MANAGER",
			expectedErr: nil,
		},
		{
			name:        "wrong_id",
			id:          "wrong_id",
			role:        "USER",
			expectedErr: errors.New("not found"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, manage_serv.UpdateRole(tc.id, tc.role))
		})
	}
}
