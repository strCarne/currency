package nbrb

import (
	"context"

	"github.com/strCarne/currency/internal/schema"
	"github.com/strCarne/currency/pkg/models"
)

const nbrbRatesRawURL = "https://api.nbrb.by/exrates/rates"

type Periodicity int

const (
	Dayly Periodicity = iota
	Monthly
)

type ParamMode int

const (
	Internal ParamMode = iota
	DigitISO4217
	AlphaISO4217
)

type Client interface {
	GetRates(
		ctx context.Context,
		onDate *models.Date,
		periodicity Periodicity,
		paramMode *ParamMode,
	) ([]schema.Rate, error)

	GetRate(
		ctx context.Context,
		curID int,
		onDate *models.Date,
		paramMode *ParamMode,
	) (*schema.Rate, error)
}
