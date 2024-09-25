package db

import (
	"database/sql"
	"fin_data_processing/internal/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"log/slog"
)

func GetDbConnection() *sql.DB {
	db, err := sql.Open("postgres", config.DSN)
	if err != nil {
		log.Fatal(err)
	}

	// Проверка подключения
	if err = db.Ping(); err != nil {
		slog.Error(err.Error())
	}

	return db
}

func GetGormDbConnection() *gorm.DB {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.DSN,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		slog.Error(err.Error())
	}

	return db
}
