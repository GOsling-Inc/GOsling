package services_test

import (
	"errors"
	"testing"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/stretchr/testify/assert"
)

func Test_Investment_CreateOrder(t *testing.T) {
	db = database.NewMock()
	inv_serv := services.NewInvestmentService(db)
	db.AddUser(&user0)
	acc := models.Account{
		Id:     "0",
		UserId: user0.Id,
		Name:   "Account",
		Type:   "Basic",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	db.AddAccount(acc)
	db.AddInvestment(models.Investment{
		Id:   911,
		Name: "Apple",
		Investors: map[string]int{
			"1111": 1,
			acc.Id: 2,
		},
	})
	orders := []models.Order{
		{
			Id:        1,
			Name:      "Яблоко",
			UserId:    user0.Id,
			AccountId: acc.Id,
			Count:     2,
			Action:    "BUY",
			Price:     200,
		},
		{
			Id:        2,
			Name:      "Apple",
			UserId:    "wrong_id",
			AccountId: acc.Id,
			Count:     2,
			Action:    "",
			Price:     200,
		},
		{
			Id:        3,
			Name:      "Apple",
			UserId:    user0.Id,
			AccountId: acc.Id,
			Count:     100,
			Action:    "BUY",
			Price:     200,
		},
		{
			Id:        4,
			Name:      "Apple",
			UserId:    user0.Id,
			AccountId: acc.Id,
			Count:     4,
			Action:    "SELL",
			Price:     200,
		},
		{
			Id:        5,
			Name:      "Apple",
			UserId:    user0.Id,
			AccountId: acc.Id,
			Count:     1,
			Action:    "BUY",
			Price:     200,
		},
		{
			Id:        6,
			Name:      "Apple",
			UserId:    user0.Id,
			AccountId: acc.Id,
			Count:     2,
			Action:    "SELL",
			Price:     200,
		},
	}
	testCases := []struct {
		name        string
		payload     models.Order
		expectedErr error
	}{
		{
			name:        "wrong_stock",
			payload:     orders[0],
			expectedErr: errors.New("wrong stock"),
		},
		{
			name:        "wrong_id",
			payload:     orders[1],
			expectedErr: errors.New("order error"),
		},
		{
			name:        "wrong_action",
			payload:     orders[1],
			expectedErr: errors.New("order error"),
		},
		{
			name:        "not_enough_money",
			payload:     orders[2],
			expectedErr: errors.New("order error"),
		},
		{
			name:        "not_enough_stocks",
			payload:     orders[3],
			expectedErr: errors.New("order error"),
		},
		{
			name:        "valid_buy",
			payload:     orders[4],
			expectedErr: nil,
		},
		{
			name:        "valid_sell",
			payload:     orders[5],
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, inv_serv.CreateOrder(tc.payload))
		})
	}
}

func Test_Investment_BuyStock(t *testing.T) {
	db = database.NewMock()
	inv_serv := services.NewInvestmentService(db)
	db.AddUser(&user0)
	basicacc := models.Account{
		Id:     "0",
		UserId: user0.Id,
		Name:   "Account",
		Type:   "Basic",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	invacc := models.Account{
		Id:     "1",
		UserId: user0.Id,
		Name:   "Account",
		Type:   "INVESTMENT",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	db.AddAccount(basicacc)
	db.AddAccount(invacc)
	db.AddInvestment(models.Investment{
		Id:   911,
		Name: "Apple",
		Investors: map[string]int{
			"1111": 1,
		},
	})
	testCases := []struct {
		name        string
		acc         models.Account
		order       models.Order
		count       int
		expectedErr error
	}{
		{
			name: "wrong_stock",
			acc:  basicacc,
			order: models.Order{
				Id:        123,
				Name:      "Яблоко",
				UserId:    user0.Id,
				AccountId: basicacc.Id,
				Count:     1,
				Action:    "BUY",
				Price:     200,
			},
			count:       1,
			expectedErr: errors.New("wrong stock"),
		},
		{
			name: "to_much_stocks",
			acc:  basicacc,
			order: models.Order{
				Id:        123,
				Name:      "Apple",
				UserId:    user0.Id,
				AccountId: basicacc.Id,
				Count:     1,
				Action:    "BUY",
				Price:     200,
			},
			count:       5,
			expectedErr: errors.New("wrong stock"),
		},
		{
			name: "wrong_acc_type",
			acc:  basicacc,
			order: models.Order{
				Id:        123,
				Name:      "Apple",
				UserId:    user0.Id,
				AccountId: basicacc.Id,
				Count:     5,
				Action:    "BUY",
				Price:     200,
			},
			count:       1,
			expectedErr: errors.New("account error"),
		},
		{
			name: "not_enough_money",
			acc:  invacc,
			order: models.Order{
				Id:        123,
				Name:      "Apple",
				UserId:    user0.Id,
				AccountId: invacc.Id,
				Count:     100,
				Action:    "BUY",
				Price:     200,
			},
			count:       10,
			expectedErr: errors.New("account error"),
		},
		{
			name: "valid_buy",
			acc:  invacc,
			order: models.Order{
				Id:        123,
				Name:      "Apple",
				UserId:    user0.Id,
				AccountId: invacc.Id,
				Count:     100,
				Action:    "BUY",
				Price:     200,
			},
			count:       5,
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, inv_serv.BuyStock(tc.acc, tc.order, tc.count))
		})
	}
}

func Test_Investment_SellStock(t *testing.T) {
	db = database.NewMock()
	inv_serv := services.NewInvestmentService(db)
	db.AddUser(&user0)
	basicacc := models.Account{
		Id:     "0",
		UserId: user0.Id,
		Name:   "Account",
		Type:   "Basic",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	invacc := models.Account{
		Id:     "1",
		UserId: user0.Id,
		Name:   "Account",
		Type:   "INVESTMENT",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	db.AddAccount(basicacc)
	db.AddAccount(invacc)
	db.AddInvestment(models.Investment{
		Id:   911,
		Name: "Apple",
		Investors: map[string]int{
			"1111":    1,
			invacc.Id: 3,
		},
	})
	testCases := []struct {
		name        string
		acc         models.Account
		order       models.Order
		count       int
		expectedErr error
	}{
		{
			name: "wrong_stock",
			acc:  basicacc,
			order: models.Order{
				Id:        123,
				Name:      "Яблоко",
				UserId:    user0.Id,
				AccountId: basicacc.Id,
				Count:     1,
				Action:    "BUY",
				Price:     200,
			},
			count:       1,
			expectedErr: errors.New("wrong stock"),
		},
		{
			name: "wrong_acc_type",
			acc:  basicacc,
			order: models.Order{
				Id:        123,
				Name:      "Apple",
				UserId:    user0.Id,
				AccountId: basicacc.Id,
				Count:     5,
				Action:    "BUY",
				Price:     200,
			},
			count:       1,
			expectedErr: errors.New("account error"),
		},
		{
			name: "valid_sell",
			acc:  invacc,
			order: models.Order{
				Id:        123,
				Name:      "Apple",
				UserId:    user0.Id,
				AccountId: invacc.Id,
				Count:     100,
				Action:    "BUY",
				Price:     200,
			},
			count:       1,
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, inv_serv.SellStock(tc.acc, tc.order, tc.count))
		})
	}
}

func Test_Invetsment_GetOrders(t *testing.T) {
	db = database.NewMock()
	inv_serv := services.NewInvestmentService(db)
	db.AddUser(&user0)
	acc := models.Account{
		Id:     "0",
		UserId: user0.Id,
		Name:   "Account",
		Type:   "INVESTMENT",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	db.AddAccount(acc)
	db.AddInvestment(models.Investment{
		Id:   911,
		Name: "Apple",
		Investors: map[string]int{
			"1111": 1,
			acc.Id: 2,
		},
	})
	orders := []models.Order{
		{
			Id:        1,
			Name:      "Apple",
			UserId:    user0.Id,
			AccountId: acc.Id,
			Count:     1,
			Action:    "BUY",
			Price:     200,
		},
		{
			Id:        2,
			Name:      "Apple",
			UserId:    user0.Id,
			AccountId: acc.Id,
			Count:     2,
			Action:    "SELL",
			Price:     200,
		},
	}
	for _, ord := range orders {
		db.CreateOrder(ord)
	}
	testCases := []struct {
		name        string
		payload     int
		order       models.Order
		expectedErr error
	}{
		{
			name:        "valid",
			payload:     1,
			order:       orders[0],
			expectedErr: nil,
		},
		{
			name:        "not_found",
			payload:     666,
			order:       models.Order{},
			expectedErr: errors.New("not found"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			o, e := inv_serv.GetOrder(tc.payload)
			assert.Equal(t, tc.order, o)
			assert.Equal(t, tc.expectedErr, e)
		})
	}
}
