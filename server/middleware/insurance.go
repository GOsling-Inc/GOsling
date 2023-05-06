package middleware

import (
	"errors"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
)

type IInsuranceMiddleware interface {
	CreateInsurance(models.Insurance) (int, error)
	GetUserInsurances(string) (int, []models.Insurance, error)
}

type InsuranceMiddleware struct {
	service services.IService
}

func NewInsuranceMiddleware(s services.IService) *InsuranceMiddleware {
	return &InsuranceMiddleware{
		service: s,
	}
}

func (m *InsuranceMiddleware) CreateInsurance(insurance models.Insurance) (int, error) {
	acc, err := m.service.GetAccountById(insurance.AccountId)
	if err != nil {
		return UNAUTHORIZED, err
	}
	if acc.UserId != insurance.UserId || acc.Amount < insurance.Part {
		return UNAUTHORIZED, errors.New("account error")
	}
	if err = m.service.CreateInsurance(insurance); err != nil {
		return INTERNAL, err
	}
	return ACCEPTED, nil
}

func (m *InsuranceMiddleware) GetUserInsurances(userId string) (int, []models.Insurance, error) {
	insurances, err := m.service.GetUserInsurances(userId)
	if err != nil {
		return INTERNAL, nil, err
	}
	return OK, insurances, nil
}
