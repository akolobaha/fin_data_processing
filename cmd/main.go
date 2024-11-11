package main

import (
	"fin_data_processing/cmd/commands"
	"fin_data_processing/internal/config"
)

const defaultEnvFilePath = ".env"

func main() {
	cfg, err := config.Parse(defaultEnvFilePath)
	if err != nil {
		panic("Ошибка парсинга конфигов")
	}
	config.InitDSN(cfg)

	commands.RunGrpc(cfg)
	commands.RunHttp(cfg)

	// Обработка сигналов завершения

}
