package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"e-store-backend/internal/repository"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger  *slog.Logger
	queries *repository.Queries
}

func main() {
	type config struct {
		addr string
		dsn  string
	}
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4210", "HTTP network address")
	flag.StringVar(&cfg.dsn, "dsn", "dev_user:dev_password@tcp(127.0.0.1:6033)/ecomm?parseTime=true", "MySQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	db, err := openDB(cfg.dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	queries := repository.New(db)

	app := &application{
		logger:  logger,
		queries: queries,
	}

	app.logger.Info("Starting server", slog.String("addr", cfg.addr))
	err = http.ListenAndServe(cfg.addr, app.routes())
	app.logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
