package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/strCarne/currency/internal/clients/rates"
	"github.com/strCarne/currency/pkg/db"
	"github.com/strCarne/currency/pkg/wrapper"
)

func GetCollectedRates(ctx echo.Context) error {
	ratesClient := rates.NewStd(db.Connection())

	rates, err := ratesClient.SelectRates(ctx.Request().Context())
	if err != nil {
		return jsonResponse(ctx, http.StatusInternalServerError, "coudn't get rates from db", err)
	}

	// return jsonResponse(ctx, http.StatusOK, "collected rates", rates)
	return wrapper.Wrap(
		"internal.routes.general.Index",
		"GetCollectedRates handler failed",
		jsonResponse(ctx, http.StatusOK, "collected rates", rates),
	)
}
