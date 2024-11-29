package service

import (
	"context"
	"encoding/json"
	"fin_data_processing/db"
	"fin_data_processing/internal/config"
	"fin_data_processing/internal/entities"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
)

func SaveFundamentalMsg(ctx context.Context, msg amqp.Delivery, cfg *config.Config) error {
	client := db.GetMongoDbConnection(ctx, cfg)
	defer client.Disconnect(ctx)

	// Get a handle for your collection
	collection := client.Database("fin").Collection("fundamentals")

	headers := msg.Headers

	data := entities.Fundamental{
		Ticker:       headers["Ticker"].(string),
		ReportMethod: headers["ReportMethod"].(string),
		PeriodType:   headers["PeriodType"].(string),
		Period:       headers["Period"].(string),
		ReportUrl:    headers["ReportUrl"].(string),
	}
	if err := json.Unmarshal(msg.Body, &data); err != nil {
		log.Printf("Ошибка при разборе сообщения: %s", err)
	}

	filter := bson.M{
		"PeriodType":   data.PeriodType,
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
