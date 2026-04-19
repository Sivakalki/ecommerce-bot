package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	XaiAPIKey   string
	DatabaseURL string
}

// Load reads the .env file and environment variables to build the Config.
func Load() (*Config, error) {
	// Optional .env loading
	_ = godotenv.Load()

	cfg := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	if cfg.DatabaseURL == "" {
		// Fallback to exactly what is defined in docker-compose if unset
		cfg.DatabaseURL = "postgres://admin:password@localhost:5433/ecommerce?sslmode=disable"
		fmt.Fprintf(os.Stderr, "Warning: DATABASE_URL not set. Falling back to default: %s\n", cfg.DatabaseURL)
	}

	return cfg, nil
}
