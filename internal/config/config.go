package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
)

type Config struct {
	ServerAddress            string `env:"SERVER_ADDRESS"`
	GrpcHost                 string `env:"GRPC_HOST"`
	GrpcPort                 string `env:"GRPC_PORT"`
	RabbitUsername           string `env:"RABBIT_USERNAME"`
	RabbitPassword           string `env:"RABBIT_PASSWORD"`
	RabbitHost               string `env:"RABBIT_HOST"`
	RabbitPort               string `env:"RABBIT_PORT"`
	RabbitQueueQuotes        string `env:"RABBIT_QUEUE_QUOTES"`
	RabbitQueueFundamentals  string `env:"RABBIT_QUEUE_FUNDAMENTALS"`
	RabbitQueueNotifications string `env:"RABBIT_QUEUE_NOTIFICATIONS"`
	MongoUsername            string `env:"MONGO_USERNAME"`
	MongoPassword            string `env:"MONGO_PASSWORD"`
	MongoHost                string `env:"MONGO_HOST"`
	MongoPort                string `env:"MONGO_PORT"`
	MongoDatabase            string `env:"MONGO_DATABASE"`
	MongoCollection          string `env:"MONGO_COLLECTION"`
	LogLevel                 string `env:"LOG_LEVEL"`
	PrometheusPort           string `env:"PROMETHEUS_PORT"`
	PrometheusHost           string `env:"PROMETHEUS_HOST"`
}

func Parse(s string) (*Config, error) {
	c := &Config{}
	if err := cleanenv.ReadConfig(s, c); err != nil {
		return nil, err
	}
	setLogLevel(c.LogLevel)
	return c, nil
}

func setLogLevel(level string) {
	switch level {
	case "debug":
		slog.SetLogLoggerLevel(-4)
	case "info":
		slog.SetLogLoggerLevel(0)
	case "warn":
		slog.SetLogLoggerLevel(4)
	case "error":
		slog.SetLogLoggerLevel(8)
	default:
		slog.SetLogLoggerLevel(4)
	}
}

func (cfg *Config) GetRabbitDSN() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%s/", cfg.RabbitUsername, cfg.RabbitPassword, cfg.RabbitHost, cfg.RabbitPort,
	)
}

func (cfg *Config) GetMongoDSN() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/admin", cfg.MongoUsername, cfg.MongoPassword, cfg.MongoHost, cfg.MongoPort)
}

func (cfg *Config) GetGrpc() string {
	return fmt.Sprintf(
		"%s:%s", cfg.GrpcHost, cfg.GrpcPort)
}

func (cfg *Config) GetPrometheusURL() string {
	return fmt.Sprintf(
		"%s:%s", cfg.PrometheusHost, cfg.PrometheusPort,
	)
}
