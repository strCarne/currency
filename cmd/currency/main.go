package main

import (
	"github.com/strCarne/currency/internal/setup"
)

func main() {
	// Setting environment variables
	setup.MustEnv()

	// Setting up GORM
	connPool := setup.MustGORM()

	// Migrations
	setup.MustMigrate(connPool)

	// Setting up logger
	logger := setup.MustLogger()

	// Starting process of enriching rates table
	enricher := setup.MustEnricher(logger, connPool)
	go enricher.Start()

	// Starting process of polling NBRB's API
	poller := setup.MustPoller(logger, enricher.RatesRx())
	go poller.Start()

	// Setting up echo server
	echoServer := setup.Echo(logger, connPool, poller)
	echoServer.Logger.Fatal(echoServer.Start(":8000"))
}
