package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDbUri  string
	MongoDbName string
	Port        string
}

// LoadConfig loads all environment variables from the .env file,
// stores them inside a Config struct,
// and returns a POINTER to that struct.
func LoadConfig() *Config {

	// Read the .env file from the project's root directory.
	//
	// Example:
	// MONGO_URI=mongodb://localhost:27017
	// PORT=8080
	//
	// If the file doesn't exist or cannot be read,
	// godotenv.Load() returns an error.
	err := godotenv.Load()

	// Check if loading the .env file failed.
	if err != nil {

		// log.Fatal() prints the error message
		// and immediately stops the application.
		//
		// We stop here because without our configuration,
		// the application doesn't know things like:
		// - Which database to connect to
		// - Which port to run on
		// - API keys
		log.Fatal("Could not load .env file:", err)
	}

	// Create a Config struct and store its ADDRESS in cfg.
	//
	// The '&' means:
	// "Create the struct and give me its memory address."
	//
	// Instead of copying the entire struct,
	// we now have one Config object in memory
	// that the whole application can share.
	cfg := &Config{

		// Read the MongoDB connection string
		// from the environment variables.
		MongoDbUri: getEnv("MONGO_URI"),

		// Read the database name.
		MongoDbName: getEnv("MONGO_DB_NAME"),

		// Read the server port.
		Port: getEnv("PORT"),
	}

	// Print a success message so we know
	// configuration loaded correctly.
	log.Println("Environment variables loaded successfully.")

	// Return the POINTER to the Config struct.
	//
	// Notice the function returns *Config.
	// That matches what cfg actually is.
	return cfg
}

func getEnv(key string) string {

	value := os.Getenv(key)

	if value == "" {
		panic(fmt.Sprintf("Environment variable %s is missing", key))
	}

	return value
}
