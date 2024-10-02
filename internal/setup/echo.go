package setup

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	slogecho "github.com/samber/slog-echo"
	"github.com/strCarne/currency/internal/controllers"
	"github.com/strCarne/currency/internal/routes"
	"gorm.io/gorm"
)

func Echo(logger *slog.Logger, connPool *gorm.DB, poller *controllers.NBRBPoller) *echo.Echo {
	echoServer := echo.New()
	echoServer.HideBanner = true

	echoServer.Use(slogecho.New(logger))

	echoServer.GET("/", routes.Index)

	ratesGroup := echoServer.Group("/rates")
	ratesGroup.GET("/all-collected", routes.Wrap(routes.AllCollectedRates, connPool, poller))
	ratesGroup.GET("/by-date/:date", routes.Wrap(routes.RatesByDate, connPool, poller))

	return echoServer
}
