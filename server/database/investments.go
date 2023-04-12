package database

import (
	"encoding/json"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/jmoiron/sqlx"
)

type InvestmentDatabase struct {
	db *sqlx.DB
}

func NewInvestmentDatabase(db *sqlx.DB) *InvestmentDatabase {
	return &InvestmentDatabase{
		db: db,
	}
}

func (d *InsuranceDatabase) GetInvestments() ([]models.Investment, error) {
	var rawInvestments []models.RawInvestment
	query := "SELECT * from  investments"

	err := d.db.Select(&rawInvestments, query)
	investments := make([]models.Investment, len(rawInvestments))
	for i, raw := range rawInvestments {
		investments[i] = models.Investment{
			Id:   raw.Id,
			Name: raw.Name,
		}
		json.Unmarshal([]byte(raw.Investors), &investments[i].Investors)
	}
	return investments, err
}
