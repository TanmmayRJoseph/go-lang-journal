package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

// to store multiple configuration values in a single struct
type Config struct {
	Port        string
	Env         string
	DatabaseURL string
}

func MustLoadConfig() Config {

	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		slog.Error("PORT environment variable is not set")
		panic("PORT environment variable is not set")
	}

	env := os.Getenv("ENV")
	if env == "" {
		slog.Error("PORT environment variable is not set")
		panic("PORT environment variable is not set")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		slog.Error("DATABASE_URL environment variable is not set")
		panic("DATABASE_URL environment variable is not set")
	}

	return Config{
		Port:        port,
		Env:         env,
		DatabaseURL: databaseURL,
	}
}
