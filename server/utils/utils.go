package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
)

var byn_usd float64
var byn_eur float64

type IUtils interface {
	MakeID() string
	Hash(string) (string, error)
	getExchanges() map[string]float64
	BYN_USD() float64
	BYN_EUR() float64
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

func (u *Utils) UpdateExchanges() {
	r, err := http.Get("https://www.nbrb.by/api/exrates/rates?periodicity=0")
	if err != nil {
		pair := models.ExchangePair{}
		content, _ := ioutil.ReadFile("env/exchanges.json")
		json.Unmarshal(content, &pair)
		byn_usd = pair.BYN_USD
		byn_eur = pair.BYN_EUR
		return
	}
	var data []map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&data)
	for _, v := range data {
		if v["Cur_Abbreviation"] == "USD" {
			byn_usd, _ = v["Cur_OfficialRate"].(float64)
		}
		if v["Cur_Abbreviation"] == "EUR" {
			byn_eur, _ = v["Cur_OfficialRate"].(float64)
		}
	}
	pair := models.ExchangePair{
		BYN_USD: byn_usd,
		BYN_EUR: byn_eur,
	}
	content, _ := json.Marshal(pair)
	err = ioutil.WriteFile("env/exchanges.json", content, 0644)
}

func BYN_USD() float64 {
	return byn_usd
}

func BYN_EUR() float64 {
	return byn_eur
}
