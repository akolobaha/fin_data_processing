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
	config.InitDbDSN(cfg)
	config.InitRabbitDSN(cfg)

	//commands.RunGrpc(cfg)
	//commands.RunHttp(cfg)
	commands.RunFundamentalsListener(cfg)

	// Обработка сигналов завершения

}
