// Package runner
package runner

import (
	"fmt"
)

type Runner struct {
	connector SSHConnector
	commands  CommandRunner
}

func New(connector SSHConnector, commands CommandRunner) *Runner {
	return &Runner{
		connector: connector,
		commands:  commands,
	}
}

func (r *Runner) Run() (string, error) {
	signer, err := r.connector.GetSigner()
	if err != nil {
		return "", fmt.Errorf("ошибка чтения ключа: %w", err)
	}

	cfg, err := r.connector.NewConfig(signer)
	if err != nil {
		return "", fmt.Errorf("ошибка создания SSH-конфигурации: %w", err)
	}

	client, err := r.connector.Connect(cfg)
	if err != nil {
		return "", fmt.Errorf("ошибка подключения к серверу: %w", err)
	}
	defer client.Close()

	fmt.Println("✅ Подключено к серверу")

	cmds, err := r.commands.LoadCommand("commands.json")
	if err != nil {
		return "", fmt.Errorf("ошибка загрузки команд: %w", err)
	}

	output, err := r.commands.RunCommand(client, cmds["wg"])
	if err != nil {
		return "", fmt.Errorf("ошибка выполнения команды: %w", err)
	}

	return output, nil
}
