package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/TanmmayRJoseph/olx-api/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		slog.Error("Please provide a valid command")
		slog.Info("Usage: migrate <up | down>")
		os.Exit(1)
	}

	cfg := config.MustLoadConfig()

	m, err := migrate.New(
		"file://migrations",
		cfg.DatabaseURL)
	if err != nil {
		slog.Error("Failed to create migrate instance", "error", err)
		os.Exit(1)
	}

	// Check the command provided by the user and accordingly run the migrations
	switch os.Args[1] {
	case "up":
		slog.Info("Running Migrations .....")
		if err := m.Up(); err != nil {
			slog.Error("Failed to run migrations", "error", err)
			os.Exit(1)
		}

	case "down":
		// PROBLEM: The `m.Down()` method rolls back all migrations, which is not always desirable. Instead, we can use `m.Steps(-1)` to roll back only the last migration.
		// slog.Info("Rolling back Migrations .....")
		// if err := m.Down(); err != nil {
		// 	slog.Error("Failed to rollback migrations", "error", err)
		// 	os.Exit(1)
		// }

		slog.Info("Rolling back Migrations .....")
		if err := m.Steps(-1); err != nil {
			slog.Error("Failed to rollback migrations", "error", err)
			os.Exit(1)
		}
	default:
		slog.Error("Invalid command provided")
		os.Exit(1)
	}

	slog.Info("Done .....")
}
