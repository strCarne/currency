package main

import (
	"github.com/strCarne/currency/internal/schema"
	"github.com/strCarne/currency/internal/setup"
	"github.com/strCarne/currency/pkg/db"
)

func main() {
	// Setting environment variables
	setup.MustEnv()

	// Setting up GORM
	setup.MustGORM()

	// Migrations
	setup.MustMigrate(db.Connection())

	// Setting up logger
	logger := setup.MustLogger()

	// Setting up channel between Poller and Enricher
	ratesFlow := make(chan []schema.Rate)
	defer close(ratesFlow)

	// Starting process of polling NBRB's API
	poller := setup.MustPoller(logger)
	go poller.Start(ratesFlow)

	// Starting process of enriching rates table
	enricher := setup.MustEnricher(logger)
	go enricher.Start(ratesFlow)

	// Setting up echo server
	echoServer := setup.Echo(logger)
	echoServer.Logger.Fatal(echoServer.Start(":8000"))
}
