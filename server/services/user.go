package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/golang-jwt/jwt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

const (
	salt = "GOsling"
)

type UserService struct {
	database *database.Database
}

func NewUserService(d *database.Database) *UserService {
	return &UserService{
		database: d,
	}
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

func (s *UserService) HashPassword(user *models.User) error {
	if len(user.Password) > 0 {
		hashed, err := s.getHashedPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashed
		log.Println(user)
		log.Println(hashed)
		return nil
	}
	return errors.New("length of password must be more than 0")
}

func (s *UserService) getHashedPassword(str string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func (s *UserService) CreateJWT(id string) (string, error) {
	claims := models.JWTClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			Id:        "user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(salt))
	if err != nil {
		return "", err
	}
	return token, nil
}