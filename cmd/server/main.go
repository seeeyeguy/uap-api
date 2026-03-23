package main

import (
	"context"
	"time"
)

func main() {
	cfg := loadConfig()
	logger := setupLogger(cfg.Environment)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := NewDB(ctx, cfg.DatabaseURL)
	if err!= nil {
		logger.Error("failed to connect to database", "error", err)
	}
	
	defer db.Close()
	logger.Info("database connected")

	srv := NewServer(cfg, logger, db)

	logger.Info("starting up",
		"port", cfg.Port,
		"enviroment", cfg.Environment,
		)

	if err := srv.Run(); err != nil {
		logger.Error("server exited with error", "error", err)
	}
}
