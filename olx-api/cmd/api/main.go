package main

import (
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/TanmmayRJoseph/olx-api/internal/config"
	"github.com/TanmmayRJoseph/olx-api/internal/database"
	"github.com/TanmmayRJoseph/olx-api/internal/handlers"
)

// For Production -> go build -o bin/main main.go -> -o stands for output

// Using MakeFile -> [package.json]
func main() {
	cfg := config.MustLoadConfig() // loading the configuration values from .env file

	// Connecting to the database using the provided database
	db, err := database.ConnectToDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	mux := http.NewServeMux() // Express.router() -> creating a new router

	// making health route
	mux.HandleFunc("GET /healthz", handlers.Health)                           // adding the health route to the router
	mux.HandleFunc("GET /listings", handlers.ListingsHandler(db))             // adding the listings route to the router
	mux.HandleFunc("DELETE /delete/{id}", handlers.DeleteListingsHandler(db)) //delete a listing

	// Configuring the server
	server := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10, // 10s
		WriteTimeout: time.Second * 10, //10s
		IdleTimeout:  time.Second * 60, //1 min
	}

	// making a simple server and logging the error if any
	slog.Info("🔥 Server running on %s", "port", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server Failed: %v", err)
	}
}
