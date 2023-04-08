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
	ProvideExchange(*models.Exchange) error
	FreezeAccount(string) error
	DeleteAccount(string) error
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

func (s *AccountService) FreezeAccount(accountId string) error {
	if err := s.database.FreezeAccount(accountId); err != nil {
		return errors.New("incorrect account")
	}
	return nil
}

func (s *AccountService) DeleteAccount(accountId string) error {
	if err := s.database.DeleteAccount(accountId); err != nil {
		return errors.New("incorrect account")
	}
	return nil
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

func (s *AccountService) ProvideExchange(exchange *models.Exchange) error {
	sender_acc, err := s.database.GetAccountById(exchange.Sender)
	if err != nil || sender_acc.Amount < exchange.SenderAmount || sender_acc.State != "ACTIVE" {
		return errors.New("sender account error, transaction cancelled")
	}
	reciever_acc, err := s.database.GetAccountById(exchange.Receiver)
	if err != nil || reciever_acc.Unit == sender_acc.Unit {
		return errors.New("reciever account error, transaction cancelled")
	}
	if sender_acc.UserId != reciever_acc.UserId {
		return errors.New("incorrect account")
	}
	if sender_acc.State != "ACTIVE" || reciever_acc.State != "ACTIVE" {
		return errors.New("one of accounts is not active")
	}
	
	if sender_acc.Unit == "BYN" && reciever_acc.Unit == "USD" {
		exchange.Course = 1 / utils.BYN_USD()
	} else if sender_acc.Unit == "BYN" && reciever_acc.Unit == "EUR" {
		exchange.Course = 1 / utils.BYN_EUR()
	} else if sender_acc.Unit == "USD" && reciever_acc.Unit == "BYN" {
		exchange.Course = utils.BYN_USD()
	} else if sender_acc.Unit == "EUR" && reciever_acc.Unit == "BYN" {
		exchange.Course = 1 / utils.BYN_USD()
	} else if sender_acc.Unit == "USD" && reciever_acc.Unit == "EUR" {
		exchange.Course = utils.BYN_USD() / utils.BYN_EUR()
	} else if sender_acc.Unit == "EUR" && reciever_acc.Unit == "USD" {
		exchange.Course = utils.BYN_EUR() / utils.BYN_USD()
	}

	exchange.ReceiverAmount = exchange.Course * exchange.SenderAmount
	if err = s.database.Exchange(exchange.Sender, exchange.Receiver, exchange.SenderAmount, exchange.ReceiverAmount); err != nil {
		return err
	}
	err = s.database.AddExchange(exchange)
	return err
}
