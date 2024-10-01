package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/strCarne/currency/pkg/models"
	"github.com/strCarne/currency/pkg/shared"
	"github.com/strCarne/currency/pkg/wrapper"
)

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

	code := http.StatusOK

	response := shared.Response[models.ProjectInfo]{
		Code:    code,
		Message: "operation completed successfully",
		Data:    info,
	}

	return wrapper.Wrap(
		"internal.routes.general.Index",
		"operation completed successfully",
		ctx.JSON(code, &response),
	)
}
