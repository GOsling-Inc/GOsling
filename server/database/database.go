package database

import (
	"log"

	"github.com/GOsling-Inc/GOsling/env"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type IAuthDatabase interface {
	GetUserByMail(string) (*models.User, error)
	GetUserById(string) (*models.User, error)
	AddUser(*models.User) error
}

type Database struct {
	IAuthDatabase
}

func New(db *sqlx.DB) *Database {
	return &Database{
		IAuthDatabase: NewAuthDatabase(db),
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
