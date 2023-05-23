package middleware_test

import (
	"testing"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/stretchr/testify/assert"
)

func Test_Investment_CreateOrder(t *testing.T) {
	db = database.NewMock()
	inv_serv := services.New(db)
	inv_mid := middleware.NewInvestmentMiddleware(inv_serv)
	db.AddUser(&user0)
	valid_acc := models.Account{
		Id:     "0",
		UserId: user0.Id,
		Name:   "Account",
		Type:   "INVESTMENT",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	invalid_acc := models.Account{
		Id:     "1",
		UserId: user0.Id,
		Name:   "Account",
		Type:   "Basic",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	}
	db.AddAccount(valid_acc)
	db.AddInvestment(models.Investment{
		Id:   911,
		Name: "Apple",
		Investors: map[string]int{
			"1111":       1,
			valid_acc.Id: 2,
		},
	})
	orders := []models.Order{
		{
			Id:        1,
			Name:      "Яблоко",
			UserId:    user0.Id,
			AccountId: valid_acc.Id,
			Count:     2,
			Action:    "BUY",
			Price:     200,
		},
		{
			Id:        2,
			Name:      "Apple",
			UserId:    "wrong_id",
			AccountId: valid_acc.Id,
			Count:     2,
			Action:    "",
			Price:     200,
		},
		{
			Id:        3,
			Name:      "Apple",
			UserId:    user0.Id,
			AccountId: valid_acc.Id,
			Count:     100,
			Action:    "BUY",
			Price:     200,
		},
		{
			Id:        4,
			Name:      "Apple",
			UserId:    user0.Id,
			AccountId: valid_acc.Id,
			Count:     4,
			Action:    "SELL",
			Price:     200,
		},
		{
			Id:        5,
			Name:      "Apple",
			UserId:    user0.Id,
			AccountId: invalid_acc.Id,
			Count:     1,
			Action:    "BUY",
			Price:     200,
		},
		{
			Id:        6,
			Name:      "Apple",
			UserId:    user0.Id,
			AccountId: valid_acc.Id,
			Count:     1,
			Action:    "BUY",
			Price:     200,
		},
		{
			Id:        7,
			Name:      "Apple",
			UserId:    user0.Id,
			AccountId: valid_acc.Id,
			Count:     2,
			Action:    "SELL",
			Price:     200,
		},
		{
			Id:        22,
			Name:      "Apple",
			UserId:    user0.Id,
			AccountId: valid_acc.Id,
			Count:     2,
			Action:    "",
			Price:     200,
		},
	}
	testCases := []struct {
		name         string
		payload      models.Order
		expectedCode int
	}{
		{
			name:         "wrong_stock",
			payload:      orders[0],
			expectedCode: 500,
		},
		{
			name:         "wrong_id",
			payload:      orders[1],
			expectedCode: 401,
		},
		{
			name:         "wrong_action",
			payload:      orders[7],
			expectedCode: 500,
		},
		{
			name:         "not_enough_money",
			payload:      orders[2],
			expectedCode: 500,
		},
		{
			name:         "not_enough_stocks",
			payload:      orders[3],
			expectedCode: 500,
		},
		{
			name:         "invalid_acc",
			payload:      orders[4],
			expectedCode: 401,
		},
		{
			name:         "valid_buy",
			payload:      orders[5],
			expectedCode: 202,
		},
		{
			name:         "valid_sell",
			payload:      orders[6],
			expectedCode: 202,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := inv_mid.CreateOrder(tc.payload)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}

func Test_Investment_BuyStock(t *testing.T) {
	db = database.NewMock()
	inv_serv := services.New(db)
	inv_mid := middleware.NewInvestmentMiddleware(inv_serv)
	db.AddUser(&user0)
	db.AddAccount(models.Account{
		Id:     "1",
		UserId: user0.Id,
		Name:   "Account",
		Type:   "INVESTMENT",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	})
	db.CreateOrder(models.Order{
		Id:        6,
		Name:      "Apple",
		UserId:    user0.Id,
		AccountId: "1",
		Count:     2,
		Action:    "SELL",
		Price:     200,
	})
	db.AddInvestment(models.Investment{
		Id:   911,
		Name: "Apple",
		Investors: map[string]int{
			"1111": 1,
			"1":    2,
		},
	})
	testCases := []struct {
		name         string
		id           string
		acc          string
		order        string
		count        string
		expectedCode int
	}{
		{
			name:         "invalid_input",
			id:           user0.Id,
			acc:          "1",
			order:        "",
			count:        "1",
			expectedCode: 500,
		},
		{
			name:         "unauthorized",
			id:           "666",
			acc:          "1",
			order:        "6",
			count:        "1",
			expectedCode: 401,
		},
		{
			name:         "wrong_order",
			id:           user0.Id,
			acc:          "1",
			order:        "7",
			count:        "1",
			expectedCode: 500,
		},
		{
			name:         "valid",
			id:           user0.Id,
			acc:          "1",
			order:        "6",
			count:        "1",
			expectedCode: 202,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := inv_mid.BuyStock(tc.order, tc.acc, tc.count, tc.id)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}

func Test_Investment_SellStock(t *testing.T) {
	db = database.NewMock()
	inv_serv := services.New(db)
	inv_nid := middleware.NewInvestmentMiddleware(inv_serv)
	db.AddUser(&user0)
	db.AddAccount(models.Account{
		Id:     "1",
		UserId: user0.Id,
		Name:   "Account",
		Type:   "INVESTMENT",
		Unit:   "BYN",
		Amount: 1234,
		State:  "ACTIVE",
	})
	db.CreateOrder(models.Order{
		Id:        6,
		Name:      "Apple",
		UserId:    user0.Id,
		AccountId: "1",
		Count:     2,
		Action:    "BUY",
		Price:     200,
	})
	db.AddInvestment(models.Investment{
		Id:   911,
		Name: "Apple",
		Investors: map[string]int{
			"1111": 1,
			"1":    2,
		},
	})
	testCases := []struct {
		name         string
		id           string
		acc          string
		order        string
		count        string
		expectedCode int
	}{
		{
			name:         "invalid_input",
			id:           user0.Id,
			acc:          "1",
			order:        "",
			count:        "1",
			expectedCode: 500,
		},
		{
			name:         "unauthorized",
			id:           "666",
			acc:          "1",
			order:        "6",
			count:        "1",
			expectedCode: 401,
		},
		{
			name:         "wrong_order",
			id:           user0.Id,
			acc:          "1",
			order:        "7",
			count:        "1",
			expectedCode: 500,
		},
		{
			name:         "valid",
			id:           user0.Id,
			acc:          "1",
			order:        "6",
			count:        "1",
			expectedCode: 202,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := inv_nid.SellStock(tc.order, tc.acc, tc.count, tc.id)
			assert.Equal(t, tc.expectedCode, c)
		})
	}
}
