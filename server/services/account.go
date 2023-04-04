package services

import (
	"errors"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/utils"
)

type IAccountService interface {
	AddAccount(*models.User, *models.Account) error
	GetUserAccounts(*models.User) ([]models.Account, error)
	ProvideTransfer(*models.Trasfer) error
}

type AccountService struct {
	database *database.Database
	Utils    *utils.Utils
}

func NewAccountService(d *database.Database, u *utils.Utils) *AccountService {
	return &AccountService{
		database: d,
		Utils:    u,
	}
}

func (s *AccountService) AddAccount(user *models.User, acc *models.Account) error {
	accs, err := s.database.GetUserAccounts(user.Id)
	if len(accs) == 12 || err != nil {
		return errors.New("can't add an account")
	}
	acc.Id = s.Utils.MakeID() + user.Id + acc.Unit
	acc.UserId = user.Id
	if err := s.database.IAccountDatabase.AddAccount(acc); err != nil {
		return err
	}
	return nil
}

func (s *AccountService) GetUserAccounts(user *models.User) ([]models.Account, error) {
	accs, err := s.database.GetUserAccounts(user.Id)
	if err != nil {
		return nil, err
	}
	return accs, nil
}

func (s *AccountService) ProvideTransfer(transfer *models.Trasfer) error {
	sender_acc, err := s.database.GetAccountById(transfer.Sender)
	if err != nil || sender_acc.Amount < transfer.Amount {
		return errors.New("sender account error, transaction cancelled")
	}
	reciever_acc, err := s.database.GetAccountById(transfer.Receiver)
	if err != nil || reciever_acc.Unit != sender_acc.Unit {
		return errors.New("reciever account error, transaction cancelled")
	}
	if sender_acc.State != "ACTIVE" || reciever_acc.State != "ACTIVE" {
		return errors.New("one of accounts is not active")
	}
	if err = s.database.Transfer(transfer.Sender, transfer.Receiver, transfer.Amount); err != nil {
		return err
	}
	err = s.database.AddTransfer(transfer)
	return err
}
