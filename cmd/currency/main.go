package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/strCarne/currency/internal/schema"
	"github.com/strCarne/currency/internal/setup"
	"github.com/strCarne/currency/pkg/db"
)

func main() {
	// Setting environment variables
	err := godotenv.Load()
	if err != nil {
		log.Panicf("failed to load .env file: %v", err)
	}

	// Migrations
	setup.MustMigrate(db.Connection())

	// Setting up logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Setting up channel between Poller and Enricher
	ratesFlow := make(chan []schema.Rate)
	defer close(ratesFlow)

	// Starting process of polling NBRB's API
	poller := setup.MustPoller(logger)
	go poller.Start(ratesFlow)

	// Setting up echo server
	echoServer := setup.Echo(logger)
	echoServer.Logger.Fatal(echoServer.Start(":8000"))
}
