package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ManageDatabase struct {
	db *sqlx.DB
}

func NewManageDatabase(db *sqlx.DB) *ManageDatabase {
	return &ManageDatabase{
		db: db,
	}
}

func (d *ManageDatabase) UpdateStatus(table, id, state string) error {
	query := fmt.Sprintf("Update %s SET state = $1 WHERE id = $2", table)
	_, err := d.db.Exec(query, state, id)
	return err
}
