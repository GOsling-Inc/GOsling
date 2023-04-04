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
}

type AccountService struct {
	database *database.Database
	Utils *utils.Utils
}

func NewAccountService(d *database.Database,  u *utils.Utils) *AccountService {
	return &AccountService{
		database: d,
		Utils: u,
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
