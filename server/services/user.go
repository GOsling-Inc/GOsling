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
	salt       = "#bXZZG0sling$$#"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
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
	log.Println(token)
	return token, nil
}

func (s *UserService) ParseJWT(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, err := token.Method.(*jwt.SigningMethodHMAC); !err {
			return nil, errors.New("invalid singing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok {
		return "", errors.New("token claims are not of type JWTClaims")
	}
	log.Println(claims.ID)
	return claims.ID, nil
}

func (s *UserService) GetUser(id string) error {
	_, err := s.database.GetUserById(id)
	if err != nil {
		return errors.New("incorrect email or password")
	}
	return nil
}
