package db

import (
	"context"
	"database/sql"
	"fin_data_processing/internal/config"
	"fmt"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"log/slog"
)

func GetDbConnection() *sql.DB {
	db, err := sql.Open("postgres", config.DbDsn)
	if err != nil {
		log.Fatal(err)
	}

	// Проверка подключения
	if err = db.Ping(); err != nil {
		slog.Error(err.Error())
	}

	return db
}

func GetMongoDbConnection(ctx context.Context, cfg *config.Config) *mongo.Client {
	clientOptions := options.Client().ApplyURI(cfg.GetMongoDSN())

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}
