package middleware

import (
	"errors"
	"strconv"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
)

type IInvestmentMiddleware interface {
	CreateOrder(models.Order) (int, error)
	Orders() (int, []models.Order)
	BuyStock(string, string, string, string) (int, error)
	SellStock(string, string, string, string) (int, error)
}

type InvestmentMiddleware struct {
	service services.IService
}

func NewInvestmentMiddleware(s services.IService) *InvestmentMiddleware {
	return &InvestmentMiddleware{
		service: s,
	}
}

func (m *InvestmentMiddleware) Orders() (int, []models.Order) {
	return OK, m.service.Orders()
}

func (m *InvestmentMiddleware) CreateOrder(order models.Order) (int, error) {
	acc, err := m.service.GetAccountById(order.AccountId)
	if err != nil {
		return UNAUTHORIZED, err
	}
	if acc.UserId != order.UserId || acc.Type != "INVESTMENT" || acc.State != "ACTIVE" {
		return UNAUTHORIZED, errors.New("account error")
	}
	if err = m.service.CreateOrder(order); err != nil {
		return INTERNAL, err
	}
	return ACCEPTED, nil
}

func (m *InvestmentMiddleware) BuyStock(OrderId, AccountId, Count, id string) (int, error) {
	if OrderId == "" || AccountId == "" || Count == "" {
		return INTERNAL, errors.New("wrong data")
	}
	acc, err := m.service.GetAccountById(AccountId)
	if err != nil || acc.UserId != id {
		return UNAUTHORIZED, err
	}
	ord_id, _ := strconv.Atoi(OrderId)
	ord_count, _ := strconv.Atoi(Count)
	ord, err := m.service.GetOrder(ord_id)
	if err != nil || ord.Action == "BUY" {
		return INTERNAL, errors.New("order error")
	}
	if err = m.service.BuyStock(acc, ord, ord_count); err != nil {
		return INTERNAL, err
	}
	return ACCEPTED, nil
}

func (m *InvestmentMiddleware) SellStock(OrderId, AccountId, Count, id string) (int, error) {
	if OrderId == "" || AccountId == "" || Count == "" {
		return INTERNAL, errors.New("wrong data")
	}
	acc, err := m.service.GetAccountById(AccountId)
	if err != nil || acc.UserId != id {
		return UNAUTHORIZED, err
	}
	ord_id, _ := strconv.Atoi(OrderId)
	ord_count, _ := strconv.Atoi(Count)
	ord, err := m.service.GetOrder(ord_id)
	if err != nil || ord.Action == "SELL" || ord.Count < ord_count {
		return INTERNAL, errors.New("order error")
	}
	if err = m.service.SellStock(acc, ord, ord_count); err != nil {
		return INTERNAL, err
	}
	return ACCEPTED, nil
}
