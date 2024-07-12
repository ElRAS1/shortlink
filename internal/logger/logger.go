package logger

import (
	"log/slog"
	"os"
)

var Logger = &slog.Logger{}

func ConfigureLogger(level int, cfg string) *slog.Logger {
	switch cfg {
	case "dev":
		Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(level)}))
	case "prod":
		Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(level)}))
	default:
		Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(level)}))
	}
	return Logger
}
