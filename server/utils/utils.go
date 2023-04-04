package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"

	"github.com/GOsling-Inc/GOsling/database"
)

type IUtils interface {
	MakeID() string
	Hash(string) (string, error)
}

type Utils struct {
	database *database.Database
}

func NewUtils(d *database.Database) *Utils {
	return &Utils{
		database: d,
	}
}

func (u *Utils) MakeID() string {
	var charset = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, 7)
	for {
		for i := range b {
			b[i] = charset[rand.Intn(len(charset))]
		}
		id := string(b)
		_, err := u.database.GetUserById(id)
		if err != nil {
			return id
		}
	}
}

func (u *Utils) Hash(str string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
