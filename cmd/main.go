package main

import (
	"context"
	"fin_data_processing/cmd/commands"
	"fin_data_processing/internal/config"
	"os"
	"os/signal"
	"syscall"
)

const defaultEnvFilePath = ".env"

func main() {
	cfg, err := config.Parse(defaultEnvFilePath)
	if err != nil {
		panic("Ошибка парсинга конфигов")
	}

	ctx, cancel := context.WithCancel(context.Background())

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
