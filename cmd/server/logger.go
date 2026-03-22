package main

import (
	"log/slog"
	"os"
)

func setupLogger(env string) *slog.Logger {
	var handler slog.Handler

	if env == "production" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	l := slog.New(handler)
	slog.SetDefault(l)
	return l
}

