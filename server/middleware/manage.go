package middleware

import (
	"errors"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
)

type IManagerMiddleware interface {
	GetConfirms() (int, []models.Unconfirmed)
	Confirm(string, string, string) (int, error)
	GetAccounts() (int, []models.Account)
	UpdateAccount(string, string) (int, error)
	GetTransactions() (int, []models.Trasfer)
	CancelTransaction(string) (int, error)
	GetUsers() (int, []models.User)
	UpdateRole(string, string) (int, error)
}

type ManagerMiddleware struct {
	service services.IService
}

func NewManagerMiddleware(s services.IService) *ManagerMiddleware {
	return &ManagerMiddleware{
		service: s,
	}
}

func (m *ManagerMiddleware) GetConfirms() (int, []models.Unconfirmed) {
	return OK, m.service.GetConfirms()
}

func (m *ManagerMiddleware) Confirm(id, table, state string) (int, error) {
	switch table {
	case "loans":
		err := m.service.ConfirmLoan(id, state)
		if err != nil {
			return INTERNAL, err
		}
		return OK, nil
	case "deposits":
		err := m.service.ConfirmDeposit(id, state)
		if err != nil {
			return INTERNAL, err
		}
		return OK, nil
	case "insurances":
		err := m.service.ConfirmInsurance(id, state)
		if err != nil {
			return INTERNAL, err
		}
		return OK, nil
	}
	return 0, nil
}

func (m *ManagerMiddleware) GetAccounts() (int, []models.Account) {
	return OK, m.service.GetAccounts()
}

func (m *ManagerMiddleware) UpdateAccount(id, state string) (int, error) {
	if _, err := m.service.GetAccountById(id); err != nil {
		return INTERNAL, err
	}
	if state != "FREEZED" && state != "BLOCKED" {
		return INTERNAL, errors.New("undefined state")
	}
	return OK, m.service.UpdateAccount(id, state)
}

func (m *ManagerMiddleware) GetTransactions() (int, []models.Trasfer) {
	return OK, m.service.GetTransactions()
}

func (m *ManagerMiddleware) CancelTransaction(id string) (int, error) {
	trs, err := m.service.GetTransferById(id)
	if err != nil {
		return INTERNAL, err
	}
	return OK, m.service.CancelTransaction(trs)
}

func (m *ManagerMiddleware) GetUsers() (int, []models.User) {
	return OK, m.service.GetUsers()
}

func (m *ManagerMiddleware) UpdateRole(id, role string) (int, error) {
	if _, err := m.service.GetUser(id); err != nil {
		return INTERNAL, err
	}
	if role != "user" && role != "manager" {
		return INTERNAL, errors.New("undefined role")
	}
	return OK, m.service.UpdateRole(id, role)
}
