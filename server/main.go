package main

import (
	"log"
	"net/http"
	"time"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/env"
	"github.com/GOsling-Inc/GOsling/handlers"
	"github.com/GOsling-Inc/GOsling/middleware"
	"github.com/GOsling-Inc/GOsling/router"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func main() {
	server := echo.New()
	server.Use(m.CORS())
	database := database.New(database.Connect())
	services := services.New(database)
	middleware := middleware.New(services)
	handlers := handlers.New(middleware)

	go func() {
		for {
			services.UpdateExchanges()
			database.UpdateLoans()
			database.UpdateDeposits()
			database.UpdateInsurances()
			time.Sleep(30 * time.Minute)
		}
	}()

	router.Init(server, handlers)

	if err := server.Start(env.GetPORT()); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
