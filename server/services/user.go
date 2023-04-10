package services

import (
	"errors"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
)

type UserService struct {
	database *database.Database
}

func NewUserService(d *database.Database) *UserService {
	return &UserService{
		database: d,
	}
}

func (s *UserService) GetUser(id string) (models.User, error) {
	user, err := s.database.GetUserById(id)
	if err != nil {
		return models.User{}, errors.New("incorrect email or password")
	}
	return user, nil
}

func (s *UserService) Change_Main_Info(u models.User) error {
	return s.database.UpdateUserData(u.Id, u.Name, u.Surname, u.Birthdate)
}

func (s *UserService) Change_Password(u models.User) error {
	return s.database.UpdatePasswordUser(u.Id, u.Password)
}
