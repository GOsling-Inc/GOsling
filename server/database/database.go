package database

import (
	"log"

	"github.com/GOsling-Inc/GOsling/env"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type IDatabase interface {
	IUserDatabase
	IAccountDatabase
	ILoanDatabase
	IDepositDatabase
	IInsuranceDatabase
	IInvestmentDatabase
	IManageDatabase
}

type Database struct {
	IUserDatabase
	IAccountDatabase
	ILoanDatabase
	IDepositDatabase
	IInsuranceDatabase
	IInvestmentDatabase
	IManageDatabase
}

func New(db *sqlx.DB) *Database {
	return &Database{
		IUserDatabase:       NewUserDatabase(db),
		IAccountDatabase:    NewAccountDatabase(db),
		ILoanDatabase:       NewLoanDatabase(db),
		IDepositDatabase:    NewDepositDatabase(db),
		IInsuranceDatabase:  NewInsuranceDatabase(db),
		IInvestmentDatabase: NewInvestmentDatabase(db),
		IManageDatabase:     NewManageDatabase(db),
	}
}

func Connect() *sqlx.DB {
	db, err := sqlx.Open("postgres", env.GetDBconfig())
	if err != nil {
		log.Fatalln("DATABASE: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("DATABASE: ", err)
	}
	return db
}
