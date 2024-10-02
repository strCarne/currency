package setup

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/strCarne/currency/internal/clients/rates"
	"github.com/strCarne/currency/internal/controllers"
	"gorm.io/gorm"
)

func MustEnricher(logger *slog.Logger, connPool *gorm.DB) *controllers.Enricher {
	attemptsNumStr := os.Getenv("ENRICH_ATTEMPTS_NUM")
	if attemptsNumStr == "" {
		panic("ENRICH_ATTEMPTS_NUM must be set")
	}

	attemptsNum, err := strconv.Atoi(attemptsNumStr)
	if err != nil {
		panic("ENRICH_ATTEMPTS_NUM must be integer")
	}

	retryDelayStr := os.Getenv("ENRICH_RETRY_DELAY")
	if retryDelayStr == "" {
		panic("ENRICH_RETRY_DELAY must be set")
	}

	retryDelay, err := time.ParseDuration(retryDelayStr)
	if err != nil {
		panic("ENRICH_RETRY_DELAY must be duration")
	}

	enricher, err := controllers.NewEnricher(rates.NewStd(connPool), logger, attemptsNum, retryDelay)
	if err != nil {
		panic(fmt.Sprintf("failed to create new NBRB poller: %v", err))
	}

	return enricher
}
