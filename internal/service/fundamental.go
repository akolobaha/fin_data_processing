package service

import (
	"context"
	"encoding/json"
	"errors"
	"fin_data_processing/db"
	"fin_data_processing/internal/config"
	"fin_data_processing/internal/entities"
	"fmt"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"strings"
)

func SaveFundamentalMsg(ctx context.Context, msg amqp.Delivery, cfg *config.Config) error {
	client := db.GetMongoDbConnection(ctx, cfg)
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			slog.Error("Failed to disconnect from database", "error", err)
		}
	}(client, ctx)

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
		slog.Error(fmt.Sprintf("Ошибка при разборе сообщения: %s", err))
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
		slog.Error(err.Error())
	}

	// Обработка результата
	if updateResult.UpsertedCount > 0 {
		slog.Info("Inserted a new document with ID:", updateResult.UpsertedID)
	} else {
		slog.Info("Updated an existing document.")
	}

	return nil
}

func GetLatestQuarterReport(ctx context.Context, cfg *config.Config, ticker string, reportMethod string) (entities.Fundamental, error) {
	client := db.GetMongoDbConnection(ctx, cfg)
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			slog.Error("MongoDB disconnect error:", "error", err)
		}
	}(client, ctx)

	collection := client.Database(cfg.MongoDatabase).Collection(cfg.MongoCollection)

	filter := bson.M{
		"ReportMethod": strings.ToUpper(reportMethod),
		"Ticker":       ticker,
		"Period":       "quarter",
	}

	opts := options.FindOne().SetSort(bson.D{
		{Key: "Report", Value: -1},
	})

	// Выполняем запрос
	var fundamental entities.Fundamental
	err := collection.FindOne(ctx, filter, opts).Decode(&fundamental)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fundamental, err
		} else {
			slog.Error(err.Error())
		}
	}

	slog.Info("Получили самый свежий отчет по эмитунту ", fundamental.Ticker, fundamental.Report)

	return fundamental, nil
}
