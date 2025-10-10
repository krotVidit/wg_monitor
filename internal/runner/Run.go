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

	sshConfig := connect.NewSSHConfig(signer, cfg)
	client, err := connect.ConnectSSH(cfg, sshConfig)
	if err != nil {
		return "", fmt.Errorf("ошибка подключения к серверу: %w", err)
	}
	defer func() {
		if e := client.Close(); e != nil {
			fmt.Println("Ошибка закрытия сессии:", e)
		}
	}()

	fmt.Println("Подключено к серверу")

	// cmd := `sudo wg show all dump | awk -v now="$(date +%s)" '$6 != 0 && (now - $6) < 180 {print $1, $4, $5, strftime("%H:%M:%S", $6)}'`
	cmd, err := command.LoadCommand("commands.json")
	if err != nil {
		return "", fmt.Errorf("ошибка загрузки файла с командами %w", err)
	}
	output, err := command.RunCommand(client, cmd["wg"])
	if err != nil {
		return "", fmt.Errorf("ошибка выполнения команды: %w", err)
	}

	return output, nil
}
