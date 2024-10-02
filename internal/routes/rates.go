//nolint:godot
package routes

//nolint:gci
import (
	"context"
	"net/http"

	"github.com/strCarne/currency/internal/clients/rates"
	"github.com/strCarne/currency/pkg/models"
	"github.com/strCarne/currency/pkg/wrapper"

	// Blank imports are needed here for swaggo to generate swagger documentation.
	// Cause: swaggo can't parse external models without it.
	_ "github.com/strCarne/currency/internal/schema"
	_ "github.com/strCarne/currency/pkg/shared"
)

// @description  Get all rates, that have been ever polled from NBRB's API and stored into the database. Can return 500
// @description  in case if DB is not reachable.
// @tags         Rates
// @summary      Returns all collected rates
// @produce      json
// @success      200  {object} shared.Response[[]schema.Rate]
// @failure      500  {object} shared.Response[error]
// @router       /rates/all-collected [get]
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

// @description  Get all rates on specified date and store them into the database.
// @tags         Rates
// @summary      Returns rates on specified date
// @produce      json
// @param        date  path  string  true  "date in format yyyy-mm-dd"
// @success      200  {object} shared.Response[[]schema.Rate]
// @failure      404  {object} shared.Response[error]
// @router       /rates/by-date/{date} [get]
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
	if err != nil || len(rates) == 0 {
		rx := ctx.Poller.Request(*date)

		rates = <-rx
		if rates == nil {
			return jsonResponse[any](ctx.EchoCtx, http.StatusNotFound, "couldn't find rates locally and remotely", nil)
		}

		fetchedFromRemote = true
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
