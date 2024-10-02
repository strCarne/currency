package routes

//nolint:gci
import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/strCarne/currency/pkg/models"
	"github.com/strCarne/currency/pkg/wrapper"

	// Blank import is needed here for swaggo to generate swagger documentation.
	// Cause: swaggo can't parse shared.Response without it.
	_ "github.com/strCarne/currency/pkg/shared"
)

// @description  Get general service info and check if service is up and running
// @tags         General
// @summary      Returns general service info
// @produce      json
// @success      200  {object} shared.Response[models.ProjectInfo]
// @router       / [get]
func Index(ctx echo.Context) error {
	info := models.ProjectInfo{
		Name: "currency",
		Description: "Service 'currency' collects all conversion rates relative" +
			" to BYN (Belorussian Rubles) from NBRB's (National Bank Of Republic of" +
			" Belarus) API once a day.",
		Version: models.Version{
			Major: 0,
			Minor: 0,
			Patch: 0,
		},
	}

	return wrapper.Wrap(
		"internal.routes.general.Index",
		"Index handler failed",
		jsonResponse(ctx, http.StatusOK, "operation completed successfully", info),
	)
}
