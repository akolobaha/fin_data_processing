package commands

import (
	"context"
	"encoding/json"
	"fin_data_processing/internal/config"
	"fin_data_processing/internal/entities"
	"fin_data_processing/internal/service"
	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
	"log"
	"log/slog"
)

func ReadFromQueue(ctx context.Context, cfg *config.Config) *cobra.Command {
	c := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {

			conn, err := amqp.Dial(cfg.GetRabbitDSN())
			if err != nil {
				slog.Error("Failed to connect to RabbitMQ: %s", err)
			}
			defer conn.Close()

			ch, err := conn.Channel()
			if err != nil {
				slog.Error("Failed to open a channel: %s", err)
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
				log.Fatalf("Failed to register a consumer: %s", err)
			}

			for {
				select {
				case msg := <-msgsFundamental:
					// Сохранение фундаментала
					log.Printf("Received fundamentals from %s: %s", cfg.RabbitQueueFundamentals, msg.Body)
					err := service.SaveFundamentalMsg(ctx, msg, cfg)
					if err != nil {
						slog.Error("Failed to save fundamentals: %s", err)
					}
				case msg := <-msgsQuotes:
					// Получили котировку
					quote := entities.Quote{}
					if err := json.Unmarshal(msg.Body, &quote); err != nil {
						log.Printf("Ошибка при разборе сообщения: %s", err)
					}

					targets := entities.FetchTargets(quote.Ticker)

					if len(targets) > 0 {
						for _, target := range targets {

							latestFundamental, err := service.GetLatestQuarterReport(ctx, cfg, quote.Ticker, target.Target.FinancialReport)
							if err != nil {
								slog.Error(err.Error())
								continue
							}

							achieved, _ := service.TargetsAchievementCheck(target, latestFundamental, quote)
							if achieved {
								entities.SetTargetAchieved(target.Target.Id, achieved)
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
