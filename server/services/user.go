package services

import (
	"errors"
	"time"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/utils"
	"github.com/dgrijalva/jwt-go"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

const (
	salt       = "Gosling"
	signingKey = "ASIJj983Jf324FJj9fj20JFsif293JFfsdf23432"
)

type IUserService interface {
	GetUser(string) (*models.User, error)
	Change_Main_Info(models.User) error
	Change_Password(models.User) error
	CreateJWT(string) (string, error)
	ParseJWT(string) (string, error)
	Validate(*models.User) error
}

type UserService struct {
	database *database.Database
	Utils *utils.Utils
}

func NewUserService(d *database.Database, u *utils.Utils) *UserService {
	return &UserService{
		database: d,
		Utils: u,
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

func (s *UserService) Validate(user *models.User) error {
	return validation.ValidateStruct(user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(8, 100)))
}