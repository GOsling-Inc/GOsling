package services

import (
	"github.com/GOsling-Inc/GOsling/models"
)

// func (m *Manager) Get() error {//хз надо или нет
//Get Manager from db
// 	return nil
// }

func (s *Service) CheckAccount(id string) (models.User, error) {
	user, err := s.database.GetUserById(id)
	return user, err
}

// func (s *Service) AcceptOperation(operation string) {//тут в зависимости от типа операции её можно подтвердить

// }

// func (s *Service) RefuseOperation(operation string) {//тут отказ в операции

// }
