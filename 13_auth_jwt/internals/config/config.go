package config

import (
	"auth_jwt/internals/utils"
	"log/slog"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string `env:"PORT" envDefault:"8080"`
	MongoURI  string `env:"MONGODB_URI" envDefault:"mongodb://localhost:27017"`
	JWTSecret string `env:"JWT_SECRET" envDefault:"secret"`
}

// Imports the .env file and returns a Config object
func LoadConfig() (*Config, error) {

	

	if err := godotenv.Load(); err != nil {
		slog.Error("Failed to load .env file", "error", err)
		return nil, err
	}

	cfg := &Config{
		Port:      utils.GetEnv("PORT", "8080"),
		MongoURI:  utils.GetEnv("MONGODB_URI", "mongodb://localhost:27017"),
		JWTSecret: utils.GetEnv("JWT_SECRET", "secret"),
	}
	return cfg, nil
}
