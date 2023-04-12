package database

import (
	"context"
	"encoding/json"
	"fmt"

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

func (d *InvestmentDatabase) GetInvestments() ([]models.Investment, error) {
	var rawInvestments []models.RawInvestment
	query := "SELECT * FROM  investments"

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

func (d *InvestmentDatabase) GetInvestment(name string) (models.Investment, error) {
	var rawInvestment models.RawInvestment
	query := "SELECT * FROM investments WHERE name = $1"

	err := d.db.Get(&rawInvestment, query, name)
	investment := models.Investment{
		Id:   rawInvestment.Id,
		Name: rawInvestment.Name,
	}

	json.Unmarshal([]byte(rawInvestment.Investors), &investment.Investors)
	return investment, err
}

func (d *InvestmentDatabase) CreateOrder(order models.Order) error {
	var id int
	query := "INSERT INTO orders (name, userid, accountId, count, action, price) values ($1, $2, $3, $4, $5, $6) RETURNING id"
	return d.db.Get(&id, query, order.Name, order.UserId, order.AccountId, order.Count, order.Action, order.Price)
}

func (d *InvestmentDatabase) GetOrders() ([]models.Order, error) {
	var orders []models.Order
	query := "SELECT * FROM  Order"

	err := d.db.Select(&orders, query)
	return orders, err
}

func (d *InvestmentDatabase) GetOrder(id int) (models.Order, error) {
	var order models.Order
	query := "SELECT * FROM  Order WHERE id = $1"

	err := d.db.Select(&order, query, id)
	return order, err
}

func (d *InvestmentDatabase) Buy(accountId string, order models.Order, count int) error {
	ctx := context.Background()
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET amount = amount - $1 WHERE id = $2", order.Price*float64(count), accountId)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET amount = amount + $1 WHERE id = $2", order.Price*float64(count), order.AccountId)
	if err != nil {
		tx.Rollback()
		return err
	}
	if order.Count == count {
		_, err = tx.ExecContext(ctx, "DELETE FROM orders WHERE id = $1", order.Id)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		_, err = tx.ExecContext(ctx, "UPDATE orders SET count = $1 WHERE id = $2", order.Count-count, order.Id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	inv, _ := d.GetInvestment(order.Name)
	if inv.Investors[order.AccountId] == count {
		_, err = tx.ExecContext(ctx, "UPDATE investments SET investors = investors::jsonb - $1::jsonb WHERE name = $2", order.AccountId, order.Name)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		investor := fmt.Sprintf("{\"%s\":%d}", order.AccountId, inv.Investors[order.AccountId]-count)
		_, err = tx.ExecContext(ctx, "UPDATE investments SET investors = investors::jsonb || $1::jsonb WHERE name = $2", investor, order.Name)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	investor := fmt.Sprintf("{\"%s\":%d}", accountId, count)
	_, err = tx.ExecContext(ctx, "UPDATE investments SET investors = investors::jsonb || $1::jsonb WHERE name = $2", investor, order.Name)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO order_operations (buyer, seller, name, count, price) values ($1, $2, $3, $4, $5)", accountId, order.AccountId, order.Name, count, order.Price)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func (d *InvestmentDatabase) Sell(accountId string, order models.Order, count int) error {
	ctx := context.Background()
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET amount = amount + $1 WHERE id = $2", order.Price*float64(count), accountId)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET amount = amount - $1 WHERE id = $2", order.Price*float64(count), order.AccountId)
	if err != nil {
		tx.Rollback()
		return err
	}
	if order.Count == count {
		_, err = tx.ExecContext(ctx, "DELETE FROM orders WHERE id = $1", order.Id)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		_, err = tx.ExecContext(ctx, "UPDATE orders SET count = $1 WHERE id = $1", order.Count-count, order.Id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	inv, _ := d.GetInvestment(order.Name)
	if inv.Investors[accountId] == count {
		_, err = tx.ExecContext(ctx, "UPDATE investments SET investors = investors::jsonb - $1 WHERE name = $2", accountId, order.Name)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		investor := fmt.Sprintf("{\"%s\":%d}", accountId, inv.Investors[accountId]-count)
		_, err = tx.ExecContext(ctx, "UPDATE investments SET investors = investors::jsonb || $1::jsonb WHERE name = $2", investor, order.Name)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	investor := fmt.Sprintf("{\"%s\":%d}", order.AccountId, inv.Investors[order.AccountId]+count)
	_, err = tx.ExecContext(ctx, "UPDATE investments SET investors = investors::jsonb || $1::jsonb WHERE name = $2", investor, order.Name)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO order_operations (buyer, seller, name, count, price) values ($1, $2, $3, $4, $5)", order.AccountId, accountId, order.Name, count, order.Price)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}
