package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"
	"time"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/dgrijalva/jwt-go"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

const (
	salt       = "Gosling"
	signingKey = "ASIJj983Jf324FJj9fj20JFsif293JFfsdf23432"
)

type UserService struct {
	database *database.Database
}

func NewUserService(d *database.Database) *UserService {
	return &UserService{
		database: d,
	}
}

func (s *UserService) CreateJWT(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ID: id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *UserService) ParseJWT(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.ID, nil
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	user, err := s.database.GetUserById(id)
	if err != nil {
		return nil, errors.New("incorrect email or password")
	}
	return user, nil
}

func (s *UserService) Change_Main_Info(u models.User) error {
	return s.database.UpdateUserData(u.Id, u.Name, u.Surname, u.Birthdate)
}

func (s *UserService) Change_Password(u models.User) error {
	return s.database.UpdatePasswordUser(u.Id, u.Password)
}

func (s *UserService) MakeID() string {
	var charset = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
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

func (s *UserService) Validate(user *models.User) error {
	return validation.ValidateStruct(user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(8, 100)))
}

func (s *UserService) Hash(str string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func (s *UserService) AddAccount(user *models.User, acc *models.Account) error {
	accs, err := s.database.GetUserAccounts(user.Id)
	if len(accs) == 12 || err != nil {
		return errors.New("can't add an account")
	}
	acc.Id = s.MakeID() + user.Id + acc.Unit
	acc.UserId = user.Id
	if err := s.database.IAccountDatabase.AddAccount(acc); err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserAccounts(user *models.User) ([]models.Account, error) {
	accs, err := s.database.GetUserAccounts(user.Id)
	if err != nil {
		return nil, err
	}
	return accs, nil
}
