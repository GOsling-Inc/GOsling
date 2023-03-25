package services

import (
	"crypto/sha256"
	"errors"
	"math/rand"

	"github.com/GOsling-Inc/GOsling/models"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func (s *Service) SignIn(u *models.User) error {
	user, err := s.database.GetUserByMail(u.Email)
	if err != nil {
		return errors.New("wrong email or password")
	}
	if u.Password != user.Password {
		return errors.New("wrong password")
	}
	return nil
}

func (s *Service) SignUp(u *models.User) error {
	user, err := s.database.GetUserByMail(u.Email)
	if err == nil {
		return errors.New("the user has already registered")
	}
	for {
		s.FormID(user)
		_, err = s.database.GetUserById(user.ID)
		if err != nil {
			break
		}
	}
	err = s.database.AddUser(u)
	return err
}

func (s *Service) FormID(u models.User) {
	var charset = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, 7)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	u.ID = string(b)
}

func (s *Service) Validate(u models.User) error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 100)))
}

func (s *Service) HashPass(u models.User) error {
	if len(u.Password) > 0 {
		hashed, err := s.getHashedPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = hashed
		return nil
	}
	return errors.New("length of password must be more than 0")
}

func (s *Service) getHashedPassword(str string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}

	return string(str), nil
}
