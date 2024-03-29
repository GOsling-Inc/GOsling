package database

import (
	"context"
	"time"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/jmoiron/sqlx"
)

type IInsuranceDatabase interface {
	AddInsurance(models.Insurance) error
	GetUserInsurances(string) ([]models.Insurance, error)
	GetInsuranceById(string) (models.Insurance, error)
	UpdateInsurances() error
}

type InsuranceDatabase struct {
	db *sqlx.DB
}

func NewInsuranceDatabase(db *sqlx.DB) *InsuranceDatabase {
	return &InsuranceDatabase{
		db: db,
	}
}

func (d *InsuranceDatabase) AddInsurance(insurance models.Insurance) error {
	var id string
	query := "INSERT INTO insurances (accountid, userid, amount, remaining, part, period, deadline) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	return d.db.Get(&id, query, insurance.AccountId, insurance.UserId, insurance.Amount, insurance.Remaining, insurance.Part, insurance.Period, insurance.Deadline)
}

func (d *InsuranceDatabase) GetUserInsurances(userId string) ([]models.Insurance, error) {
	var insurances []models.Insurance
	query := "SELECT * FROM insurances WHERE userid=$1"
	err := d.db.Select(&insurances, query, userId)
	return insurances, err
}

func (d *InsuranceDatabase) GetInsuranceById(id string) (models.Insurance, error) {
	var insurance models.Insurance
	query := "SELECT * FROM insurances WHERE id=$1"
	err := d.db.Get(&insurance, query, id)
	return insurance, err
}

func (d *InsuranceDatabase) UpdateInsurances() error {
	date := time.Now().Format("2006-01-02")
	var insurances []models.Insurance
	query := "SELECT * FROM insurances WHERE deadline=$1 AND state = $2"
	d.db.Select(&insurances, query, date, "ACTIVE")

	ctx := context.Background()
	for _, insurance := range insurances {
		tx, err := d.db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(ctx, "UPDATE accounts SET amount = amount - $1 WHERE id = $2", insurance.Part, insurance.AccountId)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = tx.ExecContext(ctx, "UPDATE insurances SET remaining = remaining + $1, deadline=deadline + INTERVAL '1 month' WHERE id = $2", insurance.Part, insurance.Id)
		if err != nil {
			tx.Rollback()
			return err
		}
		if insurance.Deadline == insurance.Period {
			_, err = tx.ExecContext(ctx, "UPDATE deposits SET state = 'CLOSED' WHERE id = $1", insurance.Id)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
	}

	return nil
}
