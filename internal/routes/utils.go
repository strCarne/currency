package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/strCarne/currency/internal/controllers"
	"github.com/strCarne/currency/pkg/shared"
	"gorm.io/gorm"
)

type Context struct {
	EchoCtx  echo.Context
	ConnPool *gorm.DB
	Poller   *controllers.NBRBPoller
}

type HandlerFunc func(Context) error

func Wrap(handler HandlerFunc, connPool *gorm.DB, poller *controllers.NBRBPoller) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return handler(Context{
			EchoCtx:  ctx,
			ConnPool: connPool,
			Poller:   poller,
		})
	}
}

func jsonResponse[T any](ctx echo.Context, code int, message string, payload T) error {
	response := shared.Response[T]{
		Code:    code,
		Message: message,
		Data:    payload,
	}

	//nolint:wrapcheck
	return ctx.JSON(code, &response)
}
