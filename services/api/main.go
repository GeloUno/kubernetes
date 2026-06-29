package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Dummy microservice for responding to HTTP calls with configurable message.
func main() {
	message, ok := os.LookupEnv("API_MESSAGE")
	if !ok {
		slog.Error("Failed to load message (environment variable API_MESSAGE not found)")
		os.Exit(1)
	}

	hostname, _ := os.Hostname()

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("%s [hostname: %s]\n", message, hostname)
		w.Write([]byte(msg))
	}))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer stop()

	serverStopped := make(chan struct{})

	go func() {
		<-ctx.Done()
		srv.Shutdown(context.Background())
		close(serverStopped)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		slog.Error("HTTP Server failed", "err", err)
		os.Exit(1)
	}

	<-serverStopped
}