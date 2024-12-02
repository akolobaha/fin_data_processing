package service

import (
	"context"
	"encoding/json"
	"errors"
	"fin_data_processing/db"
	"fin_data_processing/internal/config"
	"fin_data_processing/internal/entities"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"

	"log"
)

func SaveFundamentalMsg(ctx context.Context, msg amqp.Delivery, cfg *config.Config) error {
	client := db.GetMongoDbConnection(ctx, cfg)
	defer client.Disconnect(ctx)

	// Get a handle for your collection
	collection := client.Database(cfg.MongoDatabase).Collection(cfg.MongoCollection)

	headers := msg.Headers
	if headers == nil {
		return errors.New("no headers")
	}

	data := entities.Fundamental{
		Ticker:       headers["Ticker"].(string),
		ReportMethod: headers["ReportMethod"].(string),
		Report:       headers["Report"].(string),
		Period:       headers["Period"].(string),
		ReportUrl:    headers["ReportUrl"].(string),
		SourceUrl:    headers["SourceUrl"].(string),
	}
	if err := json.Unmarshal(msg.Body, &data); err != nil {
		log.Printf("Ошибка при разборе сообщения: %s", err)
	}

	filter := bson.M{
		"Report":       data.Report,
		"ReportMethod": data.ReportMethod,
		"Ticker":       data.Ticker,
		"Period":       data.Period,
	}

	update := bson.M{
		"$setOnInsert": data,
	}

	updateResult, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))

	if err != nil {
		log.Fatal(err)
	}

	// Обработка результата
	if updateResult.UpsertedCount > 0 {
		log.Println("Inserted a new document with ID:", updateResult.UpsertedID)
	} else {
		log.Println("Updated an existing document.")
	}

	return nil
}

func GetLatestQuarterReport(ctx context.Context, cfg *config.Config, ticker string, reportMethod string) (entities.Fundamental, error) {
	client := db.GetMongoDbConnection(ctx, cfg)
	defer client.Disconnect(ctx)

	collection := client.Database(cfg.MongoDatabase).Collection(cfg.MongoCollection)

	filter := bson.M{
		"ReportMethod": strings.ToUpper(reportMethod),
		"Ticker":       ticker,
		"Period":       "quarter",
	}

	opts := options.FindOne().SetSort(bson.D{{"Report", -1}})

	// Выполняем запрос
	var fundamental entities.Fundamental
	err := collection.FindOne(ctx, filter, opts).Decode(&fundamental)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fundamental, err
		} else {
			log.Fatal(err)
		}
	}

	return fundamental, nil
}
