package routes

import (
	"context"
	"net/http"

	"github.com/strCarne/currency/internal/clients/rates"
	"github.com/strCarne/currency/pkg/models"
	"github.com/strCarne/currency/pkg/wrapper"
)

func AllCollectedRates(ctx Context) error {
	ratesClient := rates.NewStd(ctx.ConnPool)

	rates, err := ratesClient.SelectRates(ctx.EchoCtx.Request().Context())
	if err != nil {
		return jsonResponse(ctx.EchoCtx, http.StatusInternalServerError, "coudn't get rates from db", err)
	}

	return wrapper.Wrap(
		"internal.routes.general.Index",
		"GetCollectedRates handler failed",
		jsonResponse(ctx.EchoCtx, http.StatusOK, "retrieved all collected rates", rates),
	)
}

func RatesByDate(ctx Context) error {
	dateStr := ctx.EchoCtx.Param("date")
	if dateStr == "" {
		return jsonResponse[any](ctx.EchoCtx, http.StatusBadRequest, "date is required", nil)
	}

	date, err := models.DateParse(dateStr)
	if err != nil {
		return jsonResponse[any](ctx.EchoCtx, http.StatusBadRequest, "invalid date (must be yyyy-mm-dd)", err)
	}

	ratesClient := rates.NewStd(ctx.ConnPool)

	fetchedFromRemote := false

	rates, err := ratesClient.SelectRateByDate(context.Background(), *date)
	if err != nil {
		rx := ctx.Poller.Request(*date)
		rates = <-rx
		fetchedFromRemote = true
	}

	if rates == nil {
		return jsonResponse(ctx.EchoCtx, http.StatusNotFound, "couldn't find rates locally and remotely", err)
	}

	var message string
	if fetchedFromRemote {
		message = "retrieved rates from remote API and store it in local db"
	} else {
		message = "retrieved rates from local db"
	}

	return wrapper.Wrap(
		"internal.routes.general.Index",
		"RatesByDate handler failed",
		jsonResponse(ctx.EchoCtx, http.StatusOK, message, rates),
	)
}
