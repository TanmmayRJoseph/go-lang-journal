package app

import (
	"auth_jwt/internals/config"
	"auth_jwt/internals/database"
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log/slog"
	"time"
)

type App struct {
	Config  config.Config
	MongoDB *mongo.Client   // Pointer to the MongoDB client
	DB      *mongo.Database // Pointer to the MongoDB database
}

func MyApp(ctx context.Context) (*App, error) {

	// Loading the config file
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Error loading config", "error", err)
		return nil, err
	}

	// Connecting to MongoDB
	db, err := database.ConnectToDB(ctx, cfg)
	if err != nil {
		slog.Error("Error connecting to MongoDB", "error", err)
		return nil, err
	}

	return &App{
		Config:  *cfg,
		MongoDB: db.Client,
		DB:      db.Db,
	}, nil
}

// DisconnectApp disconnects from MongoDB
func (a *App) DisconnectApp(ctx context.Context) error {

	// if mongo client missing
	if a.MongoDB == nil {
		return nil
	}

	closeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := a.MongoDB.Disconnect(closeCtx); err != nil {
		slog.Error("Error disconnecting from MongoDB", "error", err)
		return err
	}
	slog.Info("Successfully disconnected from MongoDB", "error", nil)
	return nil
}
