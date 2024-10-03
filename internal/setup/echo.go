package setup

//nolint:gci
import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	"github.com/strCarne/currency/internal/controllers"
	"github.com/strCarne/currency/internal/routes"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"

	// Blank import is used for generating Swagger documentation
	// Cause: swaggo requires it to be imported like this.
	_ "github.com/strCarne/currency/api"
)

func Echo(logger *slog.Logger, connPool *gorm.DB, poller *controllers.NBRBPoller) *echo.Echo {
	echoServer := echo.New()
	echoServer.HideBanner = true

	echoServer.Use(middleware.Recover())
	echoServer.Use(slogecho.New(logger))

	// Setting up CORS for Swagger.
	//
	//nolint:exhaustruct
	echoServer.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8000"},
		AllowMethods: []string{echo.GET},
	}))

	echoServer.GET("/", routes.Index)

	ratesGroup := echoServer.Group("/rates")
	ratesGroup.GET("/all-collected", routes.Wrap(routes.AllCollectedRates, connPool, poller))
	ratesGroup.GET("/by-date/:date", routes.Wrap(routes.RatesByDate, connPool, poller))

	echoServer.GET("/swagger/*", echoSwagger.WrapHandler)

	return echoServer
}
