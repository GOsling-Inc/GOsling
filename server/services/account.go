package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
)

var (
	byn_usd float64
	byn_eur float64
)

type IAccountService interface {
	AddAccount(string, models.Account) error
	GetAccountById(string) (models.Account, error)
	GetUserAccounts(string) ([]models.Account, error)
	DeleteAccount(string) error
	ProvideTransfer(models.Trasfer) error
	ProvideExchange(models.Exchange) error
	UpdateExchanges()
	BYN_USD() float64
	BYN_EUR() float64
}

type AccountService struct {
	database database.IDatabase
}

func NewAccountService(d database.IDatabase) *AccountService {
	return &AccountService{
		database: d,
	}
}

func (s *AccountService) AddAccount(userId string, acc models.Account) error {
	accs, err := s.database.GetUserAccounts(userId)
	if len(accs) == 12 || err != nil {
		return errors.New("can't add an account")
	}
	if err := s.database.AddAccount(acc); err != nil {
		return err
	}
	return nil
}

func (s *AccountService) GetAccountById(id string) (models.Account, error) {
	return s.database.GetAccountById(id)
}

func (s *AccountService) GetUserAccounts(userId string) ([]models.Account, error) {
	accs, err := s.database.GetUserAccounts(userId)
	if err != nil {
		return nil, err
	}
	return accs, nil
}

func (s *AccountService) DeleteAccount(accountId string) error {
	if err := s.database.DeleteAccount(accountId); err != nil {
		return errors.New("incorrect account")
	}
	return nil
}

func (s *AccountService) ProvideTransfer(transfer models.Trasfer) error {
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

func (s *AccountService) ProvideExchange(exchange models.Exchange) error {
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
		exchange.Course = 1 / s.BYN_USD()
	} else if sender_acc.Unit == "BYN" && reciever_acc.Unit == "EUR" {
		exchange.Course = 1 / s.BYN_EUR()
	} else if sender_acc.Unit == "USD" && reciever_acc.Unit == "BYN" {
		exchange.Course = s.BYN_USD()
	} else if sender_acc.Unit == "EUR" && reciever_acc.Unit == "BYN" {
		exchange.Course = 1 / s.BYN_USD()
	} else if sender_acc.Unit == "USD" && reciever_acc.Unit == "EUR" {
		exchange.Course = s.BYN_USD() / s.BYN_EUR()
	} else if sender_acc.Unit == "EUR" && reciever_acc.Unit == "USD" {
		exchange.Course = s.BYN_EUR() / s.BYN_USD()
	}

	exchange.ReceiverAmount = exchange.Course * exchange.SenderAmount
	if err = s.database.Exchange(exchange.Sender, exchange.Receiver, exchange.SenderAmount, exchange.ReceiverAmount); err != nil {
		return err
	}
	err = s.database.AddExchange(exchange)
	return err
}

func (s *AccountService) UpdateExchanges() {
	r, err := http.Get("https://www.nbrb.by/api/exrates/rates?periodicity=0")
	if err != nil {
		pair := models.ExchangePair{}
		content, _ := os.ReadFile("env/exchanges.json")
		json.Unmarshal(content, &pair)
		byn_usd = pair.BYN_USD
		byn_eur = pair.BYN_EUR
		return
	}
	var data []map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&data)
	for _, v := range data {
		if v["Cur_Abbreviation"] == "USD" {
			byn_usd, _ = v["Cur_OfficialRate"].(float64)
		}
		if v["Cur_Abbreviation"] == "EUR" {
			byn_eur, _ = v["Cur_OfficialRate"].(float64)
		}
	}
	pair := models.ExchangePair{
		BYN_USD: byn_usd,
		BYN_EUR: byn_eur,
	}
	content, _ := json.Marshal(pair)
	os.WriteFile("env/exchanges.json", content, 0644)
}

func (s *AccountService) BYN_USD() float64 {
	return byn_usd
}

func (s *AccountService) BYN_EUR() float64 {
	return byn_eur
}
