package nbrb

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/strCarne/currency/internal/models/schema"
	"github.com/strCarne/currency/pkg/models"
	"github.com/strCarne/currency/pkg/wrapper"
)

type ClientStd struct{}

func NewStd() ClientStd {
	return ClientStd{}
}

func (n ClientStd) GetRates(
	ctx context.Context,
	onDate *models.Date,
	periodicity Periodicity,
	paramMode *ParamMode,
) ([]schema.Rate, error) {
	endpoint, err := url.Parse(nbrbRatesRawURL)
	if err != nil {
		return nil, wrapper.Wrap(
			"clients.nbrb.GetRates",
			"couldn't parse NBRB rates URL",
			err,
		)
	}

	queryParams := evaluateQueryParams(onDate, &periodicity, paramMode)
	endpoint.RawQuery = queryParams.Encode()

	resp, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, wrapper.Wrap("clients.nbrb.GetRates", "couldn't get NBRB rates", err)
	}
	defer resp.Body.Close()

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapper.Wrap("clients.nbrb.GetRates", "couldn't read NBRB rates payload", err)
	}

	var rates []schema.Rate
	if err := json.Unmarshal(payload, &rates); err != nil {
		return nil, wrapper.Wrap("clients.nbrb.GetRates", "couldn't unmarshal NBRB rates payload", err)
	}

	return rates, nil
}

func (n ClientStd) GetRate(
	ctx context.Context,
	curID int,
	onDate *models.Date,
	paramMode *ParamMode,
) (*schema.Rate, error) {
	nbrbRatesByIDRawURL := nbrbRatesRawURL + "/" + strconv.Itoa(curID)

	endpoint, err := url.Parse(nbrbRatesByIDRawURL)
	if err != nil {
		return nil, wrapper.Wrap(
			"clients.nbrb.GetRates",
			"couldn't parse NBRB rates URL",
			err,
		)
	}

	queryParams := evaluateQueryParams(onDate, nil, paramMode)
	endpoint.RawQuery = queryParams.Encode()

	resp, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, wrapper.Wrap("clients.nbrb.GetRates", "couldn't get NBRB rates", err)
	}
	defer resp.Body.Close()

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapper.Wrap("clients.nbrb.GetRates", "couldn't read NBRB rate payload", err)
	}

	var rate schema.Rate
	if err := json.Unmarshal(payload, &rate); err != nil {
		return nil, wrapper.Wrap("clients.nbrb.GetRates", "couldn't unmarshal NBRB rate payload", err)
	}

	return &rate, nil
}
