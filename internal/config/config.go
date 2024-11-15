package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServerAddress           string `env:"SERVER_ADDRESS"`
	PostgresUsername        string `env:"POSTGRES_USERNAME"`
	PostgresPassword        string `env:"POSTGRES_PASSWORD"`
	PostgresHost            string `env:"POSTGRES_HOST"`
	PostgresPort            string `env:"POSTGRES_PORT"`
	PostgresDatabase        string `env:"POSTGRES_DATABASE"`
	GrpcHost                string `env:"GRPC_HOST"`
	GrpcPort                string `env:"GRPC_PORT"`
	RabbitUsername          string `env:"RABBIT_USERNAME"`
	RabbitPassword          string `env:"RABBIT_PASSWORD"`
	RabbitHost              string `env:"RABBIT_HOST"`
	RabbitPort              string `env:"RABBIT_PORT"`
	RabbitQueueQuotes       string `env:"RABBIT_QUEUE_QUOTES"`
	RabbitQueueFundamentals string `env:"RABBIT_QUEUE_FUNDAMENTALS"`
}

var DbDsn string
var RabbitDsn string

func Parse(s string) (*Config, error) {
	c := &Config{}
	if err := cleanenv.ReadConfig(s, c); err != nil {
		return nil, err
	}

	return c, nil
}

func InitDbDSN(c *Config) {
	DbDsn = fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		c.PostgresUsername, c.PostgresPassword, c.PostgresDatabase, c.PostgresHost, c.PostgresPort,
	)
}

func InitRabbitDSN(c *Config) {
	RabbitDsn = fmt.Sprintf(
		"amqp://%s:%s@%s:%s/", c.RabbitUsername, c.RabbitPassword, c.RabbitHost, c.RabbitPort,
	)
}
