//nolint:godot
package main

import (
	"github.com/strCarne/currency/internal/setup"
)

// @title Currency
// @version 1.0.0
// @description Service collects all conversion rates relative to BYN from NBRB's API once a day.
// @host 0.0.0.0:8000
// @Schemes http
// @BasePath /
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
	e := setup.Echo(logger, connPool, poller)

	// Starting server
	addr := setup.MustAddress()
	e.Logger.Fatal(e.Start(addr))
}
