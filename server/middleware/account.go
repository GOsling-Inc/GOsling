package middleware

import (
	"errors"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
)

type IAccountMiddleware interface {
	GetUserAccounts(string) (int, []models.Account, error)
	AddAccount(string, models.Account) (int, error)
	GetAccountById(string) (models.Account, error)
	DeleteAccount(string, string, string) (int, error)
	UserTransfers(id string) (int, []models.Trasfer)
	ProvideTransfer(string, models.Trasfer) (int, error)
	ProvideExchange(string, models.Exchange) (int, error)
	BYN_USD() float64
	BYN_EUR() float64
}

type AccountMiddleware struct {
	service services.IService
}

func NewAccountMiddleware(s services.IService) *AccountMiddleware {
	return &AccountMiddleware{
		service: s,
	}
}

func (a *AccountMiddleware) GetUserAccounts(id string) (int, []models.Account, error) {
	accs, err := a.service.GetUserAccounts(id)
	if err != nil {
		return INTERNAL, []models.Account{}, err
	}
	return OK, accs, nil
}

func (a *AccountMiddleware) AddAccount(id string, acc models.Account) (int, error) {
	acc.Id = a.service.MakeID() + id + acc.Unit
	acc.UserId = id
	if err := a.service.AddAccount(id, acc); err != nil {
		return INTERNAL, err
	}
	return CREATED, nil
}

func (a *AccountMiddleware) GetAccountById(userId string) (models.Account, error) {
	return a.service.GetAccountById(userId)
}

func (a *AccountMiddleware) DeleteAccount(userId, accountId, password string) (int, error) {
	password, err := a.service.Hash(password)
	if err != nil {
		return UNAUTHORIZED, err
	}
	user, err := a.service.GetUser(userId)
	if err != nil {
		return UNAUTHORIZED, err
	}
	if user.Password != password {
		return UNAUTHORIZED, errors.New("incorrect password")
	}
	_, err = a.service.GetAccountById(accountId)
	if err != nil {
		return UNAUTHORIZED, err
	}
	if err = a.service.DeleteAccount(accountId); err != nil {
		return INTERNAL, err
	}
	return ACCEPTED, nil
}

func (m *AccountMiddleware) UserTransfers(id string) (int, []models.Trasfer) {
	return OK, m.service.UserTransfers(id)
}

func (m *AccountMiddleware) ProvideTransfer(id string, transfer models.Trasfer) (int, error) {
	acc, err := m.service.GetAccountById(transfer.Sender)
	if err != nil || acc.UserId != id {
		return UNAUTHORIZED, errors.New("incorrect account")
	}
	if err = m.service.ProvideTransfer(transfer); err != nil {
		return INTERNAL, err
	}
	return ACCEPTED, nil
}

func (m *AccountMiddleware) ProvideExchange(id string, exc models.Exchange) (int, error) {
	acc, err := m.service.GetAccountById(exc.Sender)
	if err != nil || acc.UserId != id {
		return UNAUTHORIZED, errors.New("incorrect account")
	}
	acc, err = m.service.GetAccountById(exc.Receiver)
	if err != nil || acc.UserId != id {
		return UNAUTHORIZED, errors.New("incorrect account")
	}
	if err = m.service.ProvideExchange(exc); err != nil {
		return INTERNAL, err
	}
	return OK, nil
}

func (m *AccountMiddleware) BYN_USD() float64 {
	return m.service.BYN_USD()
}

func (m *AccountMiddleware) BYN_EUR() float64 {
	return m.service.BYN_EUR()
}
