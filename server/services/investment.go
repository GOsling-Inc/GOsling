package services

import (
	"errors"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
)

type IInvestmentService interface {
	Orders() []models.Order
	CreateOrder(models.Order) error
	BuyStock(models.Account, models.Order, int) error
	SellStock(models.Account, models.Order, int) error
	GetOrder(int) (models.Order, error)
}

type InvestmentService struct {
	database database.IDatabase
}

func NewInvestmentService(d database.IDatabase) *InvestmentService {
	return &InvestmentService{
		database: d,
	}
}

func (s *InvestmentService) Orders() []models.Order {
	return s.database.Orders()
}

func (s *InvestmentService) CreateOrder(order models.Order) error {
	inv, err := s.database.GetInvestment(order.Name)
	if err != nil || inv.Name == "" {
		return errors.New("wrong stock")
	}
	acc, err := s.database.GetAccountById(order.AccountId)
	if err != nil || order.Action == "" {
		return errors.New("order error")
	}
	if order.Action == "BUY" && acc.Amount < float64(order.Count)*order.Price {
		return errors.New("order error")
	}
	if order.Action == "SELL" && order.Count > inv.Investors[order.AccountId] {
		return errors.New("order error")
	}
	err = s.database.CreateOrder(order)
	return err
}

func (s *InvestmentService) BuyStock(account models.Account, order models.Order, Count int) error {
	inv, err := s.database.GetInvestment(order.Name)
	if err != nil || inv.Name == "" || Count > order.Count {
		return errors.New("wrong stock")
	}
	if account.Type != "INVESTMENT" || account.Amount < float64(Count)*order.Price {
		return errors.New("account error")
	}
	err = s.database.Buy(account.Id, order, Count)
	return err
}

func (s *InvestmentService) SellStock(account models.Account, order models.Order, Count int) error {
	inv, err := s.database.GetInvestment(order.Name)
	if err != nil || inv.Name == "" {
		return errors.New("wrong stock")
	}
	if account.Type != "INVESTMENT" {
		return errors.New("account error")
	}
	err = s.database.Sell(account.Id, order, Count)
	return err
}

func (s *InvestmentService) GetOrder(id int) (models.Order, error) {
	return s.database.GetOrder(id)
}
