package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

// Dummy microservice for testing if environment variables are set as expected
func main() {
	logLevel, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		slog.Error("Failed to load log level (environment variable LOG_LEVEL not found)")
		os.Exit(1)
	}

	if logLevel != "INFO" {
		slog.Error("Unexpected log level set (valid values: INFO)", "level", logLevel)
		os.Exit(1)
	}

	mysqlUsername, ok := os.LookupEnv("MYSQL_USERNAME")
	if !ok {
		slog.Error("Failed to load DB username (environment variable MYSQL_USERNAME not found)")
		os.Exit(1)
	}

	mysqlPassword, ok := os.LookupEnv("MYSQL_PASSWORD")
	if !ok {
		slog.Error("Failed to load DB password (environment variable  MYSQL_PASSWORD not found)")
		os.Exit(1)
	}

	if mysqlUsername != "admin" || mysqlPassword != "passw0rd" {
		slog.Error("Failed to connect to DB: incorrect credentials")
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer stop()

	slog.Info("Cooking food...")
	<-ctx.Done()
}