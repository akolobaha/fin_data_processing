package commands

import (
	"context"
	"fin_data_processing/internal/config"
	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
	"log"
	"log/slog"
)

func ReadFromQueue(cfg *config.Config, dsn string, ctx context.Context) *cobra.Command {
	c := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {

			conn, err := amqp.Dial(dsn)
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
					log.Printf("Received fundamentals from %s: %s", cfg.RabbitQueueFundamentals, msg.Body)
				case msg := <-msgsQuotes:
					log.Printf("Received quotes from %s: %s", cfg.RabbitQueueQuotes, msg.Body)
				case <-ctx.Done():
					slog.Info("Сбор данных остановлен")
					return nil
				}

			}

		},
	}

	return c
}
