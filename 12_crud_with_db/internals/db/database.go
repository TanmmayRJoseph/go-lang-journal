package database

import (
	"context"
	"log/slog"
	"time"

	"book_store/internals/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// ConnectDB creates a connection to the MongoDB server.
//
// Parameters:
//
//	cfg *config.Config
//	  - A pointer to the application's configuration.
//	  - We use a pointer because we don't want to copy the entire Config
//	    struct every time this function is called.
//	  - It also makes it possible to share the same configuration object
//	    throughout the application.
//
// Returns:
//
//	*mongo.Client
//	  - A pointer to MongoDB's client.
//	  - The client represents a pool of connections to MongoDB.
//	  - This single client is shared across the entire application.
//	  - Creating a new client for every request would be slow and waste resources.
//
//	error
//	  - If anything goes wrong while connecting or verifying the connection,
//	    an error is returned.
func ConnectDB(cfg *config.Config) (*mongo.Client, error) {

	// Create a context with a timeout.
	//
	// Context is used to control how long an operation is allowed to run.
	// Here we are saying:
	// "If connecting to MongoDB takes longer than 10 seconds,
	// automatically stop trying."
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Always release resources created by the context.
	defer cancel()

	// Create MongoDB client options.
	//
	// Here we're only setting the MongoDB URI.
	clientOptions := options.Client().
		ApplyURI(cfg.MongoDBURI)

	// Create the connection.
	//
	// At this stage the driver prepares the client and attempts
	// to establish communication with MongoDB.
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		slog.Error("Failed to connect to MongoDB", "error", err)
		return nil, err
	}

	// Ping MongoDB.
	//
	// Connect() alone doesn't guarantee the server is actually reachable.
	// Ping sends a small request to MongoDB to verify that the connection works.
	if err := client.Ping(ctx, nil); err != nil {
		slog.Error("MongoDB ping failed", "error", err)
		return nil, err
	}

	slog.Info("Successfully connected to MongoDB")

	// Return the client so the rest of the application can use it.
	return client, nil
}

func DisconnectDB(db *mongo.Client) error {
	if err := db.Disconnect(context.Background()); err != nil {
		slog.Error("Error disconnecting from MongoDB", slog.Any("error", err))
		return err
	}
	slog.Info("Successfully disconnected from MongoDB", slog.Any("error", nil))
	return nil
}
