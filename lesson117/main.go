package main

import (
	"log/slog"
	"os"
)

func main() {
	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")

}

func init() {
	env := os.Getenv("ENV")
	var handler slog.Handler
	if env == "PROD" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelWarn,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
