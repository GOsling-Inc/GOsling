package main

import (
	"log"
	"net/http"
	"time"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/env"
	"github.com/GOsling-Inc/GOsling/handlers"
	"github.com/GOsling-Inc/GOsling/router"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()
	database := database.New(database.Connect())
	services := services.New(database)
	handlers := handlers.New(services)

	go func() {
		for {
			services.UpdateExchanges()
			time.Sleep(30 * time.Minute)
		}
	} ()

	router.Init(server, handlers)

	if err := server.Start(env.GetPORT()); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
