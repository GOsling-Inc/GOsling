package services

import (
	"testing"

	"github.com/GOsling-Inc/GOsling/models"
)

func TestUser(t *testing.T) *models.User {
	return &models.User{
		Email:    "user@example.org",
		Password: "password",
	}
}
