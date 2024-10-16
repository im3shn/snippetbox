package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"im3shn/snippetbox/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP netowork address")

	dsn := flag.String(
		"dsn",
		"web:pass@tcp(127.0.0.1:3306)/snippetbox?parseTime=true",
		"MySQL data source name",
	)
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		logger: logger,
		snippets: &models.SnippetModel{
			DB: db,
		},
	}

	app.logger.Info("starting server", slog.Any("addr", *addr))

	err = http.ListenAndServe(*addr, app.routes())
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
