package middleware

import (
	"errors"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
)

type UserMiddleware struct {
	service *services.Service
}

func NewUserMiddleware(s *services.Service) *UserMiddleware {
	return &UserMiddleware{
		service: s,
	}
}

func (m *UserMiddleware) GetUser(id string) (int, models.User, error) {
	user, err := m.service.GetUser(id)
	if err != nil {
		return UNAUTHORIZED, user, err
	} else {
		return OK, user, err
	}
}

func (m *UserMiddleware) Change_Main_Info(user models.User) (int, error) {
	if err := m.service.Change_Main_Info(user); err != nil {
		return UNAUTHORIZED, err
	} 
	return ACCEPTED, nil
}

func (m *UserMiddleware) Change_Password(id, oldPassword, newPassword string) (int, error) {
	oldPassword, _ = m.service.Hash(oldPassword)
	newPassword, _ = m.service.Hash(newPassword)
	if (len(newPassword) < 8 || len(newPassword) > 100) {
		return UNAUTHORIZED, errors.New("incorrect size of password")
	}
	user, err := m.service.GetUser(id)
	if err != nil {
		return UNAUTHORIZED, err
	}
	if user.Password != oldPassword {
		return UNAUTHORIZED, errors.New("passwords dont not match")
	}
	user.Password = newPassword
	
	m.service.Change_Password(user)
	return ACCEPTED, nil
}
