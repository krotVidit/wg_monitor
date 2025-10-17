package main

import (
	"fmt"
	"log"

	"wg-monitor/app/internal/command"
	"wg-monitor/app/internal/connect"
	"wg-monitor/app/internal/domain"
	"wg-monitor/app/internal/runner"
)

func main() {
	cfg, err := domain.LoadConfig("config.json")
	if err != nil {
		log.Fatal("ошибка чтения конфигурации:", err)
	}
	connector := &connect.SSHService{Cfg: cfg}
	commandRunner := &command.CommandService{}

	r := runner.New(connector, commandRunner)

	if output, err := r.Run(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(output)
	}
}
