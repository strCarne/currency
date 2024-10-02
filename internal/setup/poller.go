package setup

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/strCarne/currency/internal/clients/nbrb"
	"github.com/strCarne/currency/internal/controllers"
	"github.com/strCarne/currency/internal/schema"
)

func MustPoller(logger *slog.Logger, enricherRx chan<- []schema.Rate) *controllers.NBRBPoller {
	attemptsNumStr := os.Getenv("POLL_ATTEMPTS_NUM")
	if attemptsNumStr == "" {
		panic("POLL_ATTEMPTS_NUM must be set")
	}

	attemptsNum, err := strconv.Atoi(attemptsNumStr)
	if err != nil {
		panic("POLL_ATTEMPTS_NUM must be integer")
	}

	retryDelayStr := os.Getenv("POLL_RETRY_DELAY")
	if retryDelayStr == "" {
		panic("POLL_RETRY_DELAY must be set")
	}

	retryDelay, err := time.ParseDuration(retryDelayStr)
	if err != nil {
		panic("POLL_RETRY_DELAY must be duration")
	}

	poller, err := controllers.NewNBRBPoller(nbrb.NewStd(), logger, attemptsNum, retryDelay, enricherRx)
	if err != nil {
		panic(fmt.Sprintf("failed to create new NBRB poller: %v", err))
	}

	return poller
}
