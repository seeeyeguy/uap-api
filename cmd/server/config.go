// cmd/server/main.go

package main

import (
	"log"
	"os"
)

type Config struct {
	Port	 		string
	DatabaseURL		string
	RedisURL		string
	ClerkSecretKey		string
	AnthropicAPIKey		string
	ChromaURL		string
	Environment		string
}

func loadConfig() Config {	
	cfg := Config{
		Port: 			getEnv("PORT", "8080"),
		DatabaseURL:		requireEnv("DATABASE_URL"),
		RedisURL:		requireEnv("REDIS_URL"),
		ClerkSecretKey: 	requireEnv("CLERK_SECRET_KEY"),
		AnthropicAPIKey:	requireEnv("ANTHROPIC_API_KEY"),
		ChromaURL:		getEnv("CHROMA_URL", "http://localhost:8000"),
		Environment:		getEnv("ENVIRONMENT", "development"),
	}
	return cfg
}

// getEnv returns fallback if a required variable isn't set
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

// requireEnv crashes immediately if a required variable is missing
// better to fail {loud at startup than silently fail later
func requireEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("required environment variable %s is not set", key)
	}
	return val
}
