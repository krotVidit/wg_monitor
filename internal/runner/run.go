// Package runner
package runner

import (
	"fmt"
	"sort"
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
	defer func() {
		if e := client.Close(); e != nil {
			fmt.Println("Ошибка закрытия сессии:", e)
		}
	}()

	fmt.Println("✅ Подключено к серверу")

	cmds, err := r.commands.LoadCommand("commands.json")
	if err != nil {
		return "", fmt.Errorf("ошибка загрузки команд: %w", err)
	}

	key, err := selectCommand(cmds)
	if err != nil {
		return "", fmt.Errorf("ошибка выбора комманды: %w", err)
	}

	output, err := r.commands.RunCommand(client, cmds[key])
	if err != nil {
		return "", fmt.Errorf("ошибка выполнения команды: %w", err)
	}

	return output, nil
}

func selectCommand(cmds map[string]string) (string, error) {
	fmt.Println("Доступные команды")
	keys := make([]string, 0, len(cmds))

	for name := range cmds {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	for i, name := range keys {
		fmt.Printf("  %d. %s\n", i+1, name)
	}

	var selectedUser int
	fmt.Print("Ввыбор команды: ")
	_, err := fmt.Scan(&selectedUser)
	if err != nil || selectedUser < 1 || selectedUser > len(keys) {
		return "", fmt.Errorf("некорректный выбор")
	}

	selectedKey := keys[selectedUser-1]
	return selectedKey, err
}
