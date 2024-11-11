package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServerAddress    string `env:"SERVER_ADDRESS"`
	PostgresUsername string `env:"POSTGRES_USERNAME"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     string `env:"POSTGRES_PORT"`
	PostgresDatabase string `env:"POSTGRES_DATABASE"`
	GrpcHost         string `env:"GRPC_HOST"`
	GrpcPort         string `env:"GRPC_PORT"`
}

var Cfg *Config

var DSN string

func Parse(s string) (*Config, error) {
	c := &Config{}
	if err := cleanenv.ReadConfig(s, c); err != nil {
		return nil, err
	}

	Cfg = c
	return c, nil
}

func InitDSN(c *Config) {
	DSN = fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		c.PostgresUsername, c.PostgresPassword, c.PostgresDatabase, c.PostgresHost, c.PostgresPort,
	)
}
