package database

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectToDatabase(databaseURL string) (*sql.DB, error) {

	// Step 1: Open a connection pool to the database.
	// Note: sql.Open() doesn't actually connect yet — it just sets things up
	// and validates the arguments. The real connection happens lazily,
	// which is why we "ping" it below to make sure it actually works.
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		slog.Error("Error connecting to Database")
		return nil, fmt.Errorf("Error connecting to Database: %w", err)
	}

	// Create a context that automatically "times out" after 5 seconds.
	// This prevents our program from hanging forever if the database
	// is slow to respond or unreachable.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	// "defer" means this line runs right before the function returns.
	// Calling cancel() here frees up resources tied to the context
	// once we're done with it, so nothing leaks.
	defer cancel()

	// Step 2: Ping the database to confirm we can actually talk to it.
	// This is what triggers the real network connection and checks
	// that the credentials/URL are correct.
	if err = db.PingContext(ctx); err != nil {
		slog.Error("Error pinging the database")
		return nil, fmt.Errorf("Error pinging the database: %w", err)
	}

	// Configure the connection pool so we don't overwhelm the database
	// with too many simultaneous connections, and so idle connections
	// get recycled instead of going stale.
	db.SetMaxOpenConns(25) // Max number of connections open at the same time
	db.SetMaxIdleConns(25) // Max number of unused ("idle") connections kept ready to reuse

	// How long a single connection can be reused before it's closed and
	// replaced with a fresh one. This helps avoid issues with connections
	// that have gone stale or hit network/load balancer timeouts.
	db.SetConnMaxLifetime(5 * time.Minute) // 5 minutes

	// Step 3: If we got this far, the connection is good to go.
	slog.Info("✅ Successfully connected to the database")

	// Step 4: Hand back the ready-to-use database connection pool.
	return db, nil
}
