package main

import (
	"auth_jwt/internals/app"
	"auth_jwt/internals/httpserver"
	"context"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	slog.Info("Welcome in the class of Auth with JWT")

	// root context for the whole server
	ctx := context.Background()

	// Load the main app logic file
	a, err := app.MyApp(ctx)
	if err != nil {
		slog.Error("Error creating the app", "error", err)
		return
	}
	// Close the main app logic file
	defer func() {
		if err := a.DisconnectApp(ctx); err != nil {
			slog.Error("Error disconnecting from MongoDB", "error", err)
		}
	}()

	// ===================================================================================
	// registering all the routes in the main file
	router := httpserver.ServerRoutes(a)

	// classic Go way to start a web server
	server := &http.Server{
		Addr:              ":8080",         // listen on port 8080
		Handler:           router,          // use the router as the handler
		ReadHeaderTimeout: 5 * time.Second, // timeout for reading headers
	}

	// listen on port 8080
	slog.Info("Server is listening on port 8080")
	err = server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			slog.Info("Server closed")
			return
		}
		slog.Error("Error starting the server", "error", err)
	}

}
