package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
)

type IAuthService interface {
	SignIn(user *models.User) error
	SignUp(user *models.User) error
	MakeID() string
	Hash(str string) (string, error)
}

type AuthService struct {
	database database.IDatabase
}

func NewAuthService(d database.IDatabase) *AuthService {
	return &AuthService{
		database: d,
	}
}

func (s *AuthService) SignIn(user *models.User) error {
	tempUser, err := s.database.GetUserByMail(user.Email)
	if err != nil {
		return errors.New("incorrect email")
	}
	if user.Password != tempUser.Password {
		return errors.New("incorrect password")
	}
	user.Id = tempUser.Id
	return nil
}

func (s *AuthService) SignUp(user *models.User) error {
	_, err := s.database.GetUserByMail(user.Email)
	if err == nil {
		return errors.New("user with this email already registered")
	}
	err = s.database.AddUser(user)
	return err
}

func (s *AuthService) MakeID() string {
	charset := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, 7)
	for {
		for i := range b {
			b[i] = charset[rand.Intn(len(charset))]
		}
		id := string(b)
		_, err := s.database.GetUserById(id)
		if err != nil {
			return id
		}
	}
}

func (s *AuthService) Hash(str string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
