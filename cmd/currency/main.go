package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	slogecho "github.com/samber/slog-echo"
	"github.com/strCarne/currency/internal/routes"
	"github.com/strCarne/currency/pkg/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("failed to load .env file: %v", err)
	}

	db.InitializeGORM()

	echoServer := echo.New()
	echoServer.HideBanner = true

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	echoServer.Use(slogecho.New(logger))

	echoServer.GET("/", routes.Index)

	echoServer.Logger.Fatal(echoServer.Start(":8000"))
}
