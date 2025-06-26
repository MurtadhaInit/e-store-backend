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
	type config struct {
		port int
	}
	var cfg config
	flag.IntVar(&cfg.port, "port", 4210, "Server port")
	flag.Parse()

	addr := fmt.Sprintf(":%d", cfg.port)

	dsn, dsn_set := os.LookupEnv("DSN")
	if !dsn_set || dsn == "" {
		// fmt.Println("DSN environment variable not set. Could not connect to the database.")
		// os.Exit(1)
		panic("DSN environment variable not set. Could not connect to the database.")
	}

	app, err := app.NewApplication(dsn)
	if err != nil {
		panic(err)
	}
	defer app.DB.Close()

	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         addr,
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Info("Starting server", slog.String("addr", addr))

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Error(err.Error())
		os.Exit(1)
	}
}
