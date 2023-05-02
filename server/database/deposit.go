package database

import (
	"context"
	"time"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/jmoiron/sqlx"
)

type DepositDatabase struct {
	db *sqlx.DB
}

func NewDepositDatabase(db *sqlx.DB) *DepositDatabase {
	return &DepositDatabase{
		db: db,
	}
}

func (d *DepositDatabase) AddDeposit(deposit models.Deposit) error {
	query := "INSERT INTO deposits (accountid, userid, amount, remaining, part, percent, period, deadline) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err := d.db.Exec(query, deposit.AccountId, deposit.UserId, deposit.Amount, deposit.Remaining, deposit.Part, deposit.Percent, deposit.Period, deposit.Deadline)
	return err
}

func (d *DepositDatabase) GetUserDeposits(userId string) ([]models.Deposit, error) {
	var deposits []models.Deposit
	query := "SELECT * FROM deposits WHERE userid=$1"
	err := d.db.Select(&deposits, query, userId)
	return deposits, err
}

func (d *DepositDatabase) UpdateDeposits() error {
	date := time.Now().Format("2006-01-02")
	var deposits []models.Deposit
	query := "SELECT * FROM deposits WHERE deadline = $1 AND state = $2"
	d.db.Select(&deposits, query, date, "ACTIVE")

	ctx := context.Background()
	for _, deposit := range deposits {
		tx, err := d.db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(ctx, "UPDATE deposits SET remaining = remaining + $1, deadline=deadline + INTERVAL '1 month' WHERE id = $2", deposit.Part, deposit.Id)
		if err != nil {
			tx.Rollback()
			return err
		}
		if deposit.Deadline == deposit.Period {
			_, err = tx.ExecContext(ctx, "UPDATE deposits SET state = 'CLOSED' WHERE id = $1", deposit.Id)
			if err != nil {
				tx.Rollback()
				return err
			}
			_, err = tx.ExecContext(ctx, "UPDATE accounts SET amount = amount + $1 WHERE id = $2", deposit.Remaining+deposit.Part, deposit.Id)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
	}

	return nil
}
