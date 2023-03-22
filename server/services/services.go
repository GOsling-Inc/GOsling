package services

import (
	"github.com/GOsling-Inc/GOsling/database"
)

type Service struct {
	database *database.Database
}

func New(d *database.Database) *Service {
	return &Service{
		database: d,
	}
}
