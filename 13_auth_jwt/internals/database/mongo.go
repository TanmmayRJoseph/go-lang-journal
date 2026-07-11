package database

import (
	"auth_jwt/internals/config"
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
	Db     *mongo.Database
}

func ConnectToDB(ctx context.Context, cfg *config.Config) (*Mongo, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	clientOpt := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(clientOpt)
	if err != nil {
		slog.Error("Error connecting to MongoDB", "error", err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		slog.Error("Error pinging MongoDB", "error", err)
		return nil, err
	}

	slog.Info("Successfully connected to MongoDB")

	return &Mongo{
		Client: client,
		Db:     client.Database("book_store"),
	}, nil
}
