package controllers

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/strCarne/currency/internal/clients/rates"
	"github.com/strCarne/currency/internal/schema"
)

var (
	ErrNewEnricher = errors.New("failed to create new enricher")
	ErrEnrich      = errors.New("failed to enrich")
)

type Enricher struct {
	client      rates.DBClientInsert
	logger      *slog.Logger
	attemptsNum int
	retryDelay  time.Duration
}

func NewEnricher(
	client rates.DBClientInsert,
	logger *slog.Logger,
	attemptsNum int,
	retryDelay time.Duration,
) (*Enricher, error) {
	if client == nil {
		return nil, ErrNewEnricher
	}

	if logger == nil {
		return nil, ErrNewEnricher
	}

	if attemptsNum <= 0 {
		return nil, ErrNewEnricher
	}

	if retryDelay.Milliseconds() < 0 {
		return nil, ErrNewEnricher
	}

	return &Enricher{
		client:      client,
		logger:      logger,
		attemptsNum: attemptsNum,
		retryDelay:  retryDelay,
	}, nil
}

func (e Enricher) Start(source <-chan []schema.Rate) {
	e.EnrichFromBackup()

	for rates := range source {
		e.EnrichFromBackup()

		err := e.Enrich(rates)
		if err != nil {
			e.logger.Error("enriching failed", slog.String("cause", err.Error()))
			e.Backup(rates)
		}
	}
}

func (e Enricher) Enrich(rates []schema.Rate) error {
	for attempt := range e.attemptsNum {
		e.logger.Info("enrich attempt", slog.Int("number", attempt+1))

		err := e.InstantEnrich(rates)

		if err == nil {
			e.logger.Info("enriching completed successfully", slog.Any("data", rates))

			return nil
		}

		var mysqlError *mysql.MySQLError
		if errors.As(err, &mysqlError) && mysqlError.Number == 1062 {
			e.logger.Info("enriching completed successfully", slog.String("cause", "data already exists"))

			return nil
		}

		e.logger.Info("enriching failed", slog.String("cause", err.Error()), slog.Int("attempt", attempt+1))

		time.Sleep(e.retryDelay)
	}

	return ErrEnrich
}

func (e Enricher) InstantEnrich(rates []schema.Rate) error {
	//nolint:wrapcheck
	return e.client.InsertRates(context.Background(), rates)
}

func (e Enricher) Backup(_ []schema.Rate) {
}

func (e Enricher) EnrichFromBackup() {}
