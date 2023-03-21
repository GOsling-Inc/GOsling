package env

import "fmt"

const (
	PORT        string = ":1337"
	DB_HOST     string = "localhost"
	DB_PORT     string = "5432"
	DB_USER     string = "postgres"
	DB_PASSWORD string = "qwerty"
	DB_NAME     string = "gosling"
)

func GetPORT() string {
	return PORT
}

func GetDBconfig() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
}
