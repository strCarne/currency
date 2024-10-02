package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/strCarne/currency/pkg/shared"
)

func jsonResponse[T any](ctx echo.Context, code int, message string, payload T) error {
	response := shared.Response[T]{
		Code:    code,
		Message: message,
		Data:    payload,
	}

	//nolint:wrapcheck
	return ctx.JSON(code, &response)
}
