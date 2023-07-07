package services

import (
	"strconv"
	"time"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
)

type IInsuranceService interface {
	CreateInsurance(models.Insurance) error
	GetUserInsurances(string) ([]models.Insurance, error)
}

type InsuranceService struct {
	database database.IDatabase
}

func NewInsuranceService(d database.IDatabase) *InsuranceService {
	return &InsuranceService{
		database: d,
	}
}

func (s *InsuranceService) CreateInsurance(insurance models.Insurance) error {
	insurance.Remaining = 0
	insurance.Deadline = time.Now().AddDate(0, 0, 30).Format("2006-01-02")
	per, _ := strconv.Atoi(insurance.Period)
	insurance.Period = time.Now().AddDate(per, 0, 0).Format("2006-01-02")
	insurance.Part = insurance.Amount / float64(per*12) * 5 / 100
	err := s.database.AddInsurance(insurance)
	return err
}

func (s *InsuranceService) GetUserInsurances(id string) ([]models.Insurance, error) {
	return s.database.GetUserInsurances(id)
}
