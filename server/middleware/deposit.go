package middleware

import (
	"errors"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
)

type IDepositMiddleware interface {
	CreateDeposit(models.Deposit) (int, error)
	GetUserDeposits(string) (int, []models.Deposit, error)
}

type DepositMiddleware struct {
	service services.IService
}

func NewDepositMiddleware(s services.IService) *DepositMiddleware {
	return &DepositMiddleware{
		service: s,
	}
}

func (m *DepositMiddleware) CreateDeposit(deposit models.Deposit) (int, error) {
	acc, err := m.service.GetAccountById(deposit.AccountId)
	if err != nil {
		return UNAUTHORIZED, err
	}
	if acc.UserId != deposit.UserId || acc.Amount < deposit.Amount {
		return UNAUTHORIZED, errors.New("account error")
	}
	if err = m.service.CreateDeposit(deposit); err != nil {
		return INTERNAL, err
	}
	return ACCEPTED, nil
}

func (m *DepositMiddleware) GetUserDeposits(userId string) (int, []models.Deposit, error) {
	deposits, err := m.service.GetUserDeposits(userId)
	if err != nil {
		return INTERNAL, nil, err
	}
	return OK, deposits, nil
}
