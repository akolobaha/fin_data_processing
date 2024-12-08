package commands

import (
	"context"
	"encoding/json"
	"fin_data_processing/internal/config"
	"fin_data_processing/internal/entities"
	"fin_data_processing/internal/service"
	"fin_data_processing/internal/transport"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
	"log/slog"
)

func ReadFromQueue(ctx context.Context, cfg *config.Config) *cobra.Command {
	c := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
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
				true, // auto-ack
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
				true, // auto-ack
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
					slog.Error(fmt.Sprintf("Received fundamentals from %s: %s", cfg.RabbitQueueFundamentals, msg.Body))
					err := service.SaveFundamentalMsg(ctx, msg, cfg)
					if err != nil {
						slog.Error("Failed to save fundamentals: %s", "error", err)
					}
				case msg := <-msgsQuotes:
					// Получили котировку
					quote := entities.Quote{}
					if err := json.Unmarshal(msg.Body, &quote); err != nil {
						slog.Error(fmt.Sprintf("Ошибка при разборе сообщения: %s", err))
					}

					targets := entities.FetchTargets(quote.Ticker, cfg)

					if len(targets) > 0 {
						for _, target := range targets {
							latestFundamental, err := service.GetLatestQuarterReport(ctx, cfg, quote.Ticker, target.Target.FinancialReport)
							if err != nil {
								slog.Error(err.Error())
								continue
							}

							achieved, resultValue, err := service.TargetsAchievementCheck(target, latestFundamental, quote)
							if err != nil {
								slog.Error(err.Error())
								continue
							}
							if achieved {
								target.ResultValue = resultValue
								entities.SetTargetAchieved(target.Target.Id, achieved)
								service.SendNotificationMessage(target, rabbit, cfg.RabbitQueueNotifications)
							}
						}
					}

				case <-ctx.Done():
					slog.Info("Сервис обработки данных остановлен")
					return nil
				}

			}

		},
	}

	return c
}
