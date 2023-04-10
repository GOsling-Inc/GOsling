package services

import (
	"strconv"
	"time"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
)

type DepositService struct {
	database *database.Database
}

func NewDepositService(d *database.Database) *DepositService {
	return &DepositService{
		database: d,
	}
}

func (s *DepositService) CreateDeposit(deposit models.Deposit) error {
	deposit.Remaining = 0
	deposit.Part = (deposit.Amount / 12) * deposit.Percent / 100
	deposit.Deadline = time.Now().AddDate(0, 0, 30).Format("2006-01-02")
	per, _ := strconv.Atoi(deposit.Period)
	deposit.Period = time.Now().AddDate(per, 0, 0).Format("2006-01-02")
	err := s.database.AddDeposit(deposit)
	return err
}

func (s *DepositService) GetUserDeposits(id string) ([]models.Deposit, error) {
	return s.database.GetUserDeposits(id)
}
