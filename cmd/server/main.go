package main

func main() {
	cfg := loadConfig()
	logger := setupLogger(cfg.Environment)

	logger.Info("starting up",
		"port", cfg.Port,
		"enviroment", cfg.Environment,
		)
}
