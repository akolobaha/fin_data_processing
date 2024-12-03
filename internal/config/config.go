package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServerAddress            string `env:"SERVER_ADDRESS"`
	PostgresUsername         string `env:"POSTGRES_USERNAME"`
	PostgresPassword         string `env:"POSTGRES_PASSWORD"`
	PostgresHost             string `env:"POSTGRES_HOST"`
	PostgresPort             string `env:"POSTGRES_PORT"`
	PostgresDatabase         string `env:"POSTGRES_DATABASE"`
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
}

func Parse(s string) (*Config, error) {
	c := &Config{}
	if err := cleanenv.ReadConfig(s, c); err != nil {
		return nil, err
	}

	return c, nil
}

func (cfg *Config) GetPostgresDSN() string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.PostgresUsername, cfg.PostgresPassword, cfg.PostgresDatabase, cfg.PostgresHost, cfg.PostgresPort,
	)
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
