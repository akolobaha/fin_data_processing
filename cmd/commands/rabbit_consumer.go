package commands

import (
	"context"
	"encoding/json"
	"fin_data_processing/internal/config"
	"fin_data_processing/internal/entities"
	"fin_data_processing/internal/monitoring"
	"fin_data_processing/internal/service"
	"fin_data_processing/internal/transport"
	"fmt"
	"github.com/streadway/amqp"
	"log/slog"
)

func ReadFromQueue(ctx context.Context, cfg *config.Config) error {
	rabbit := transport.New()
	rabbit.InitConn(cfg)
	defer rabbit.ConnClose()

	conn, err := amqp.Dial(cfg.GetRabbitDSN())
	if err != nil {
		slog.Error("Failed to connect to RabbitMQ: %s", "error", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		slog.Error("Failed to open a channel: %s", "error", err)
	}
	defer ch.Close()

	msgsFundamental, err := ch.Consume(
		cfg.RabbitQueueFundamentals,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		slog.Error("Failed to register a consumer: ", "error", err)
	}

	msgsQuotes, err := ch.Consume(
		cfg.RabbitQueueQuotes,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		slog.Error("Failed to register a consumer: ", "error", err)
	}

	for {
		select {
		case msg := <-msgsFundamental:
			// Сохранение фундаментала
			slog.Info(fmt.Sprintf("Received fundamentals from %s: %s", cfg.RabbitQueueFundamentals, msg.Body))
			// TODO: записать также в кэш
			err := service.SaveFundamentalMsg(ctx, msg, cfg)
			if err != nil {
				monitoring.ProcessingErrorCount.WithLabelValues(fmt.Sprintf("Failed to save fundamentals: %s", err)).Inc()
				slog.Error("Failed to save fundamentals: %s", "error", err)
			}
			err = msg.Ack(false)
			if err != nil {
				monitoring.ProcessingErrorCount.WithLabelValues(err.Error()).Inc()
				slog.Error(err.Error())
				return err
			}
			monitoring.ProcessingSuccessCount.WithLabelValues("Фундаментальные данные успешно сохранены").Inc()
		case msg := <-msgsQuotes:
			// Получили котировку
			quote := entities.Quote{}
			if err := json.Unmarshal(msg.Body, &quote); err != nil {
				monitoring.ProcessingErrorCount.WithLabelValues(err.Error()).Inc()
				slog.Error(fmt.Sprintf("Ошибка при разборе сообщения с целью: %s", err))
			}

			targets := entities.FetchTargets(quote.Ticker, cfg)

			if len(targets) > 0 {
				for _, target := range targets {
					// TODO: Получить последний фундаментал из кэша
					latestFundamental, err := service.GetLatestQuarterReport(ctx, cfg, quote.Ticker, target.Target.FinancialReport)
					if err != nil {
						monitoring.ProcessingErrorCount.WithLabelValues(err.Error()).Inc()
						slog.Error(err.Error())
						continue
					}

					achieved, resultValue, err := service.TargetsAchievementCheck(target, latestFundamental, quote)
					if err != nil {
						monitoring.ProcessingErrorCount.WithLabelValues(err.Error()).Inc()
						slog.Error(err.Error())
						continue
					}
					if achieved {
						monitoring.ProcessingSuccessCount.WithLabelValues("Цель достигнута").Inc()
						target.ResultValue = resultValue
						entities.SetTargetAchieved(target.Target.ID, achieved)
						service.SendNotificationMessage(target, rabbit, cfg.RabbitQueueNotifications)
					}
				}
			}

			//err := msg.Ack(false)
			//if err != nil {
			//	slog.Error(err.Error())
			//	return err
			//}

		case <-ctx.Done():
			slog.Info("Сервис обработки данных остановлен")
			return nil
		}

	}

}
