package services

import (
	"errors"

	"github.com/GOsling-Inc/GOsling/models"
)

func (s *Service) SignIn(u *models.User) error {
	user, err := s.database.GetUserByMail(u.Email)
	if err != nil {
		return errors.New("incorrect email or password")
	}
	if u.Password != user.Password {
		return errors.New("incorrect email or password")
	}
	return nil
}

func (s *Service) SignUp(u *models.User) error {
	user, err := s.database.GetUserByMail(u.Email)
	if err == nil {
		return errors.New("user with this email already registered")
	}
	for {
		u.FormID()
		_, err = s.database.GetUserById(user.ID)
		if err != nil {
			break
		}
	}
	err = s.database.AddUser(u)
	return err
}
