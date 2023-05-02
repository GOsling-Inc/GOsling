package services

import "github.com/GOsling-Inc/GOsling/database"

type ManagerService struct {
	database *database.Database
}

func NewManagerService(d *database.Database) *ManagerService {
	return &ManagerService{
		database: d,
	}
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
