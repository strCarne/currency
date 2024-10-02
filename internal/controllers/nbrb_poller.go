package controllers

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/strCarne/currency/internal/clients/nbrb"
	"github.com/strCarne/currency/internal/schema"
	"github.com/strCarne/currency/pkg/models"
)

type NBRBPoller struct {
	client      nbrb.Client
	logger      *slog.Logger
	attemptsNum int
	retryDelay  time.Duration
	enricherRx  chan<- []schema.Rate
}

var (
	ErrNewNBRBPoller = errors.New("failed to create new NBRB poller")
	ErrPoll          = errors.New("failed to poll NBRB")
)

func NewNBRBPoller(
	client nbrb.Client,
	logger *slog.Logger,
	attemptsNum int,
	retryDelay time.Duration,
	enricherRx chan<- []schema.Rate,
) (*NBRBPoller, error) {
	if client == nil {
		return nil, ErrNewNBRBPoller
	}

	if logger == nil {
		return nil, ErrNewNBRBPoller
	}

	if attemptsNum <= 0 {
		return nil, ErrNewNBRBPoller
	}

	if retryDelay.Milliseconds() < 0 {
		return nil, ErrNewNBRBPoller
	}

	return &NBRBPoller{
		client:      client,
		logger:      logger,
		attemptsNum: attemptsNum,
		retryDelay:  retryDelay,
		enricherRx:  enricherRx,
	}, nil
}

func (n NBRBPoller) Start() {
	n.logger.Info("started polling at %v", slog.Time("timestamp", time.Now()))
	n.PollAndTransmit(n.currentDate(), n.enricherRx)

	for {
		n.sleepUntilNextDay()
		n.PollAndTransmit(n.currentDate(), n.enricherRx)
	}
}

func (n NBRBPoller) currentDate() *models.Date {
	loc, err := time.LoadLocation("Europe/Minsk")
	if err != nil {
		panic("failed to load location")
	}

	currentTime := time.Now().In(loc)

	return &models.Date{Year: currentTime.Year(), Month: int(currentTime.Month()), Day: currentTime.Day()}
}

func (n NBRBPoller) sleepUntilNextDay() {
	loc, err := time.LoadLocation("Europe/Minsk")
	if err != nil {
		panic("failed to load location")
	}

	now := time.Now().In(loc)
	nextDay := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 1, 0, 0, now.Location())
	untilNextDay := time.Until(nextDay)
	time.Sleep(untilNextDay)
}

func (n NBRBPoller) PollAndTransmit(date *models.Date, destination chan<- []schema.Rate) {
	rates, err := n.InstantPoll(date)
	if err == nil {
		destination <- rates
	} else {
		n.logger.Error("failed to poll", slog.String("cause", err.Error()))
		destination <- nil
	}
}

func (n NBRBPoller) Poll(date *models.Date) ([]schema.Rate, error) {
	for attempt := range n.attemptsNum {
		n.logger.Info("poll attempt", slog.Int("number", attempt+1))

		rates, err := n.InstantPoll(date)
		if err == nil {
			n.logger.Info("polling completed successfully", slog.Any("data", rates))

			return rates, nil
		}

		n.logger.Info("polling failed", slog.String("cause", err.Error()), slog.Int("attempt", attempt+1))

		time.Sleep(n.retryDelay)
	}

	return nil, ErrPoll
}

func (n NBRBPoller) InstantPoll(date *models.Date) ([]schema.Rate, error) {
	rates, err := n.client.GetRates(context.Background(), date, nbrb.Dayly, nil)
	if err == nil {
		return rates, nil
	}

	return nil, ErrPoll
}

func (n NBRBPoller) Request(date models.Date) <-chan []schema.Rate {
	n.logger.Info("new request", slog.Any("requested date", date))

	ratesTx := make(chan []schema.Rate)

	go func() {
		rates, err := n.InstantPoll(&date)
		if err != nil {
			n.logger.Error("failed to poll", slog.String("cause", err.Error()))
			ratesTx <- nil
		} else {
			ratesTx <- rates
			n.enricherRx <- rates
		}
	}()

	return ratesTx
}
