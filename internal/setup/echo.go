package setup

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	slogecho "github.com/samber/slog-echo"
	"github.com/strCarne/currency/internal/routes"
)

func Echo(logger *slog.Logger) *echo.Echo {
	echoServer := echo.New()
	echoServer.HideBanner = true

	echoServer.Use(slogecho.New(logger))

	echoServer.GET("/", routes.Index)

	return echoServer
}
