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

	app, err := app.NewApplication()
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
