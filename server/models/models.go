package models

import (
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	Role      string `json:"role"`
	Birthdate string `json:"birthdate"`
}

type Account struct {
	Id     string  `json:"id"`
	UserId string  `json:"userid"`
	Name   string  `json:"name"`
	Type   string  `json:"type"` // BASIC / BUSINESS / DEPOSIT
	Unit   string  `json:"unit"` // BYN / USD / EUR
	Amount float64 `json:"amount"`
	State  string  `json:"state"` // ACTIVE / BLOCKED / FROZEN
}

type Trasfer struct {
	Id       string  `json:"id"`
	Receiver string  `json:"receiver"`
	Sender   string  `json:"sender"`
	Amount   float64 `json:"amount"`
}

type Exchange struct {
	Id             string  `json:"id"`
	Receiver       string  `json:"receiver"`
	Sender         string  `json:"sender"`
	ReceiverAmount float64 `json:"receiveramount"`
	SenderAmount   float64 `json:"senderamount"`
	Course         float64 `json:"course"`
}

type Loan struct {
	Id        string  `json:"id"`
	AccountId string  `json:"accountId"`
	UserId    string  `json:"userid"`
	Amount    float64 `json:"amount"`
	Remaining float64 `json:"remaining"`
	Part      float64 `json:"part"`
	Percent   float64 `json:"percent"`
	Period    string  `json:"period"`
	Deadline  string  `json:"deadline"`
	State     string  `json:"state"` // ACTIVE / CLOSED
}

type Deposit struct {
	Id        string  `json:"id"`
	AccountId string  `json:"accountId"`
	UserId    string  `json:"userid"`
	Amount    float64 `json:"amount"`
	Remaining float64 `json:"remaining"`
	Part      float64 `json:"part"`
	Percent   float64 `json:"percent"`
	Period    string  `json:"period"`
	Deadline  string  `json:"deadline"`
	State     string  `json:"state"` // ACTIVE / CLOSED
}

type Insurance struct {
	Id        string  `json:"id"`
	AccountId string  `json:"accountId"`
	UserId    string  `json:"userid"`
	Amount    float64 `json:"amount"`
	Remaining float64 `json:"remaining"`
	Part      float64 `json:"part"`
	Period    string  `json:"period"`
	Deadline  string  `json:"deadline"`
	State     string  `json:"state"` // ACTIVE / CLOSED
}

type RawInvestment struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Investors string `json:"investors"`
}

type Investment struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	Investors map[string]int `json:"investors"`
}

type InvestOrder struct {
	Name   string  `json:"name"`
	UserId string  `json:"userid"`
	Count  int     `json:"count"`
	Action string  `json:"action"` // BUY / SELL
	Price  float64 `json:"price"`
}

type ExchangePair struct {
	BYN_USD float64 `json:"BYN_USD"`
	BYN_EUR float64 `json:"BYN_EUR"`
}

type JWTClaims struct {
	jwt.StandardClaims
	ID string `json:"id"`
}
