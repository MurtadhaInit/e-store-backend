package main

import (
	"e-store-backend/internal/app"
	"e-store-backend/internal/routes"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "startup failed: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	var port int
	flag.IntVar(&port, "port", 4210, "Server port")
	flag.Parse()

	addr := fmt.Sprintf(":%d", port)

	application, err := app.NewApplication()
	if err != nil {
		return err
	}
	defer func() { _ = application.DB.Close() }()

	server := &http.Server{
		Addr:         addr,
		Handler:      routes.SetupRoutes(application),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	application.Logger.Info("Starting server", slog.String("addr", addr))
	return server.ListenAndServe()
}
