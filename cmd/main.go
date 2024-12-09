package main

import (
	"context"
	"fin_data_processing/cmd/commands"
	"fin_data_processing/internal/config"
	"fin_data_processing/internal/monitoring"
	"os"
	"os/signal"
	"syscall"
)

const defaultEnvFilePath = ".env"

func init() {
	monitoring.RegisterPrometheus()
}

func main() {
	cfg, err := config.Parse(defaultEnvFilePath)
	if err != nil {
		panic("Ошибка парсинга конфигов")
	}

	ctx, cancel := context.WithCancel(context.Background())
	monitoring.RunPrometheusServer(cfg.GetPrometheusURL())

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
		<-exit
		cancel()
	}()

	err = commands.ReadFromQueue(ctx, cfg)
	if err != nil {
		return
	}
}
