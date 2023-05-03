package services

import (
	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
)

type ManagerService struct {
	database *database.Database
}

func NewManagerService(d *database.Database) *ManagerService {
	return &ManagerService{
		database: d,
	}
}

func (s *ManagerService) GetConfirms() []models.Unconfirmed {
	return s.database.GetUnconfirmed()
}

func (s *ManagerService) ConfirmLoan(id, state string) error {
	loan, err := s.database.GetLoanById(id)
	if err != nil {
		return err
	}
	loan.State = state
	return s.database.ConfirmLoan(loan)
}

func (s *ManagerService) ConfirmDeposit(id, state string) error {
	deposit, err := s.database.GetDepositById(id)
	if err != nil {
		return err
	}
	deposit.State = state
	return s.database.ConfirmDeposit(deposit)
}

func (s *ManagerService) ConfirmInsurance(id, state string) error {
	insurance, err := s.database.GetInsuranceById(id)
	if err != nil {
		return err
	}
	insurance.State = state
	return s.database.ConfirmInsurance(insurance)
}

func (s *ManagerService) GetAccounts() []models.Account {
	return s.database.GetAccounts()
}

func (s *ManagerService) UpdateAccount(id, state string) error {
	return s.database.UpdateAccount(id, state)
}

func (s *ManagerService) GetTransactions() []models.Trasfer {
	return s.database.GetTransactions()
}

func (s *ManagerService) GetTransferById(id string) (models.Trasfer, error) {
	return s.database.GetTransferById(id)
}

func (s *ManagerService) CancelTransaction(trs models.Trasfer) error {
	return s.database.CancelTransaction(trs)
}

func (s *ManagerService) GetUsers() []models.User {
	return s.database.GetUsers()
}

func (s *ManagerService) UpdateUser(id, role string) error {
	return s.database.UpdateUser(id, role)
}
