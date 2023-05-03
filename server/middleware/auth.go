package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

const (
	salt       = "Gosling"
	signingKey = "ASIJj983Jf324FJj9fj20JFsif293JFfsdf23432"
)

var sessions map[string]string = make(map[string]string)

type AuthMiddleware struct {
	service *services.Service
}

func NewAuthMiddleware(s *services.Service) *AuthMiddleware {
	return &AuthMiddleware{
		service: s,
	}
}

func (m *AuthMiddleware) SignIn(user *models.User) (int, error) {
	if err := m.Validate(*user); err != nil {
		return UNAUTHORIZED, err
	}
	user.Password, _ = m.service.Hash(user.Password)
	if err := m.service.SignIn(user); err != nil {
		return UNAUTHORIZED, err
	}
	return OK, nil
}

func (m *AuthMiddleware) SignUp(user *models.User) (int, error) {
	if err := m.Validate(*user); err != nil {
		return UNAUTHORIZED, err
	}
	user.Id = m.service.MakeID()
	user.Password, _ = m.service.Hash(user.Password)
	if err := m.service.SignUp(user); err != nil {
		return UNAUTHORIZED, err
	}
	return CREATED, nil
}

func (m *AuthMiddleware) CreateJWT(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ID: id,
	})
	signedToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	sessions[id] = signedToken
	return signedToken, nil
}

func (m *AuthMiddleware) Auth(header http.Header) string {
	id := m.parseJWT(header)
	user, err := m.service.GetUser(id)
	if err != nil {
		return ""
	}

	session, ok := sessions[user.Id]
	if !ok {
		sessions[user.Id] = header["Token"][0]
	} else {
		if session != header["Token"][0] {
			return ""
		}
	}
	return id
}

func (m *AuthMiddleware) AuthManager(id string) error {
	user, err := m.service.GetUser(id)
	if err != nil {
		return errors.New("incorrect id")
	}
	if user.Role == "user" {
		return errors.New("access denied")
	}
	return nil
}

func (m *AuthMiddleware) AuthAdmin(id string) error {
	user, err := m.service.GetUser(id)
	if err != nil {
		return errors.New("incorrect id")
	}
	if user.Role != "admin" {
		return errors.New("access denied")
	}
	return nil
}

func (m *AuthMiddleware) parseJWT(header http.Header) string {
	if header["Token"] == nil {
		return ""
	}
	id, err := m.parseToken(header["Token"][0])
	if err != nil {
		return ""
	}
	return id
}

func (m *AuthMiddleware) parseToken(accessToken string) (string, error) {
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

func (m *AuthMiddleware) Validate(user models.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(8, 100)),
	)
}

func (m *AuthMiddleware) DBTEST() error {
	return m.service.DBTEST()
}
