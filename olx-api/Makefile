# Makefile for building and running the Go application

# .PHONY target to avoid conflicts with files named 'build' or 'run'
.PHONY: build run

build:
	@go build -o bin/api ./cmd/api

# Run the application after building it
run: build
	@./bin/api



# for database migration commands
build-migrate:
	@go build -o bin/migrate ./cmd/migrate

migrate-up: build-migrate
	@go run ./cmd/migrate up

migrate-down:
	@go run ./cmd/migrate down