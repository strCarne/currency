package controllers

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/strCarne/currency/internal/clients/nbrb"
	"github.com/strCarne/currency/internal/schema"
)

type NBRBPoller struct {
	client      nbrb.Client
	logger      *slog.Logger
	attemptsNum int
	retryDelay  time.Duration
}

var (
	ErrNewNBRBPoller = errors.New("failed to create new NBRB poller")
	ErrPoll          = errors.New("failed to poll NBRB")
)

const Day = 24 * time.Hour

func NewNBRBPoller(
	client nbrb.Client,
	logger *slog.Logger,
	attemptsNum int,
	retryDelay time.Duration,
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
	}, nil
}

func (n NBRBPoller) Start(destination chan<- []schema.Rate) {
	n.logger.Info("started polling at %v", slog.Time("timestamp", time.Now()))
	n.PollAndTransmit(destination)

	for {
		n.sleepUntilNextDay()
		n.PollAndTransmit(destination)
	}
}

func (n NBRBPoller) sleepUntilNextDay() {
	now := time.Now()
	nextDay := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 1, 0, 0, now.Location())
	untilNextDay := time.Until(nextDay)
	time.Sleep(untilNextDay)
}

func (n NBRBPoller) PollAndTransmit(destination chan<- []schema.Rate) {
	rates, err := n.Poll()
	if err == nil {
		destination <- rates
	} else {
		n.logger.Error("failed to start polling", slog.String("cause", err.Error()))
	}
}

func (n NBRBPoller) Poll() ([]schema.Rate, error) {
	for attempt := range n.attemptsNum {
		n.logger.Info("poll attempt", slog.Int("number", attempt+1))

		rates, err := n.InstantPoll()
		if err == nil {
			n.logger.Info("polling completed successfully", slog.Any("data", rates))

			return rates, nil
		}

		n.logger.Info("polling failed", slog.String("cause", err.Error()), slog.Int("attempt", attempt+1))

		time.Sleep(n.retryDelay)
	}

	return nil, ErrPoll
}

func (n NBRBPoller) InstantPoll() ([]schema.Rate, error) {
	rates, err := n.client.GetRates(context.Background(), nil, nbrb.Dayly, nil)
	if err == nil {
		return rates, nil
	}

	return nil, ErrPoll
}
