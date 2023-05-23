package services

import (
	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
)

type IManageService interface {
	GetConfirms() []models.Unconfirmed
	ConfirmLoan(string, string) error
	ConfirmDeposit(string, string) error
	ConfirmInsurance(string, string) error
	GetAccounts() []models.Account
	UpdateAccount(string, string) error
	GetTransactions() []models.Trasfer
	GetTransferById(string) (models.Trasfer, error)
	CancelTransaction(models.Trasfer) error
	GetUsers() []models.User
	UpdateRole(string, string) error
}

type ManageService struct {
	database database.IDatabase
}

func NewManageService(d database.IDatabase) *ManageService {
	return &ManageService{
		database: d,
	}
}

func (s *ManageService) GetConfirms() []models.Unconfirmed {
	return s.database.GetUnconfirmed()
}

func (s *ManageService) ConfirmLoan(id, state string) error {
	loan, err := s.database.GetLoanById(id)
	if err != nil {
		return err
	}
	loan.State = state
	return s.database.ConfirmLoan(loan)
}

func (s *ManageService) ConfirmDeposit(id, state string) error {
	deposit, err := s.database.GetDepositById(id)
	if err != nil {
		return err
	}
	deposit.State = state
	return s.database.ConfirmDeposit(deposit)
}

func (s *ManageService) ConfirmInsurance(id, state string) error {
	insurance, err := s.database.GetInsuranceById(id)
	if err != nil {
		return err
	}
	insurance.State = state
	return s.database.ConfirmInsurance(insurance)
}

func (s *ManageService) GetAccounts() []models.Account {
	return s.database.GetAccounts()
}

func (s *ManageService) UpdateAccount(id, state string) error {
	return s.database.UpdateAccount(id, state)
}

func (s *ManageService) GetTransactions() []models.Trasfer {
	return s.database.GetTransactions()
}

func (s *ManageService) GetTransferById(id string) (models.Trasfer, error) {
	return s.database.GetTransferById(id)
}

func (s *ManageService) CancelTransaction(trs models.Trasfer) error {
	return s.database.CancelTransaction(trs)
}

func (s *ManageService) GetUsers() []models.User {
	return s.database.GetUsers()
}

func (s *ManageService) UpdateRole(id, role string) error {
	return s.database.UpdateRole(id, role)
}
