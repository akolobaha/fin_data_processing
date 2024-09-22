package db

import (
	"database/sql"
	"fin_data_processing/internal/config"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("postgres", config.ConnString)
	if err != nil {
		log.Fatal(err)
	}

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		slog.Error(err.Error())
	}

	return db
}
