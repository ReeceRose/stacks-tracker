package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ReeceRose/stacks-tracker/internal/server"
)

var (
	Version        = "dev"
	CommitHash     = "n/a"
	BuildTimestamp = "n/a"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info(fmt.Sprintf("Version: %s", Version))
	logger.Info(fmt.Sprintf("Commit Hash: %s", CommitHash))
	logger.Info(fmt.Sprintf("Build Timestamp: %s", BuildTimestamp))

	api := server.New(logger)
	api.Start("8080")
	defer api.Shutdown()
}
