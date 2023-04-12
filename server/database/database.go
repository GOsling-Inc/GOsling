package database

import (
	"log"

	"github.com/GOsling-Inc/GOsling/env"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	*UserDatabase
	*AccountDatabase
	*LoanDatabase
	*DepositDatabase
	*InsuranceDatabase
	*InvestmentDatabase
}

func New(db *sqlx.DB) *Database {
	return &Database{
		UserDatabase:    NewUserDatabase(db),
		AccountDatabase: NewAccountDatabase(db),
		LoanDatabase:    NewLoanDatabase(db),
		DepositDatabase: NewDepositDatabase(db),
		InsuranceDatabase: NewInsuranceDatabase(db),
		InvestmentDatabase: NewInvestmentDatabase(db),
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
