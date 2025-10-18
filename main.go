package main

import (
	"fmt"
	"log"

	"wg-monitor/app/internal/command"
	"wg-monitor/app/internal/connect"
	"wg-monitor/app/internal/domain"
	"wg-monitor/app/internal/runner"
	"wg-monitor/app/internal/ui"
)

func main() {
	cli := &ui.CLI{}
	cfg, err := domain.LoadConfig("config.json")
	if err != nil {
		log.Fatal("ошибка чтения конфигурации:", err)
	}
	connector := &connect.SSHService{Cfg: cfg}
	commandRunner := &command.CommandService{}

	r := runner.New(connector, commandRunner, cli)

	if output, err := r.Run(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(output)
	}
}
