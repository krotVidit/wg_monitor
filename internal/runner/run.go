// Package runner
package runner

import (
	"fmt"

	"wg-monitor/app/internal/command"
	"wg-monitor/app/internal/connect"
	"wg-monitor/app/internal/domain"
)

func Run() (string, error) {
	cfg, err := domain.LoadConfig("config.json")
	if err != nil {
		return "", fmt.Errorf("ошибка чтения файла json: %w", err)
	}

	signer, err := connect.GetSigner(cfg)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения ключа: %w", err)
	}
	sshConfig, err := connect.NewSSHConfig(signer, cfg)
	if err != nil {
		return "", fmt.Errorf("ошибка создания SSH-конфигурации: %w", err)
	}

	client, err := connect.ConnectSSH(cfg, sshConfig)
	if err != nil {
		return "", fmt.Errorf("ошибка подключения к серверу: %w", err)
	}
	defer func() {
		if e := client.Close(); e != nil {
			fmt.Println("Ошибка закрытия сессии:", e)
		}
	}()

	fmt.Println("✅ Подключено к серверу")

	cmd, err := command.LoadCommand("commands.json")
	if err != nil {
		return "", fmt.Errorf("ошибка загрузки файла с командами %w", err)
	}
	output, err := command.RunCommand(client, cmd["wg"])
	if err != nil {
		return "", fmt.Errorf("❌ ошибка выполнения команды: %w", err)
	}

	return output, nil
}
