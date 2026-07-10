package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

// Config stores all the application's configuration values.
type Config struct {
	Port       string `env:"PORT" envDefault:"7000"`
	MongoDBURI string `env:"MONGODB_URI" envDefault:"mongodb://localhost:27017"`
	DbUsername string `env:"DB_USERNAME" envDefault:"root"`
	DbPassword string `env:"DB_PASSWORD" envDefault:"root"`
}

// LoadConfig loads the .env file and returns a Config object.
func LoadConfig() (*Config, error) {

	// Load environment variables from the .env file.
	if err := godotenv.Load(); err != nil {
		slog.Error("Failed to load .env file", "error", err)
		return nil, err
	}

	// Create the config using helper functions &Config means we are passing the address of the config struct
	cfg := &Config{
		Port:       getEnv("PORT", "7000"),
		MongoDBURI: getRequiredEnv("MONGODB_URI"),
		DbUsername: getEnv("DB_USERNAME", "root"),
		DbPassword: getEnv("DB_PASSWORD", "root"),
	}

	return cfg, nil
}

// getEnv returns the environment variable if it exists.
// Otherwise it returns the default value.
func getEnv(key, defaultValue string) string {

	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

// getRequiredEnv returns the environment variable.
// It panics if the variable is missing because the application
// cannot run without required configuration.
func getRequiredEnv(key string) string {

	value := os.Getenv(key)

	if value == "" {
		panic(fmt.Sprintf("required environment variable %s is missing", key))
	}

	return value
}
