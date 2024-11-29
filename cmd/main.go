package main

import (
	"context"
	"fin_data_processing/cmd/commands"
	"fin_data_processing/internal/config"
	"fmt"
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

	cmd := commands.ReadFromQueue(ctx, cfg)

	//secs := entities.FetchSecurities()
	//targ := entities.FetchTargets("ROSN")
	//targ := entities.SetTargetAchieved(9, false)
	//
	//fmt.Println(secs, ctx.Err())
	//fmt.Println(targ, ctx.Err())

	//select {}
	err = cmd.ExecuteContext(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
}
