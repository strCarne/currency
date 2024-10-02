package setup

import (
	"log/slog"
	"os"
)

func MustLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	return logger
}
