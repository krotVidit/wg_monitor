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

	for {
		key, err := selectCommand(cmds)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		if key == "exit" {
			fmt.Println("👋 Выход из программы.")
			break
		}

		output, err := r.commands.RunCommand(client, cmds[key])
		if err != nil {
			fmt.Println("Ошибка выполнения команды:", err)
			continue
		}

		fmt.Printf("\n===== Результат команды '%s' =====\n%s\n", key, output)
	}

	return "Завершено.", nil
}

func selectCommand(cmds map[string]string) (string, error) {
	fmt.Println("\nДоступные команды:")
	keys := make([]string, 0, len(cmds))
	for name := range cmds {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	for i, name := range keys {
		fmt.Printf("  %d. %s\n", i+1, name)
	}
	fmt.Printf("  0. Выйти\n")

	var selectedUser int
	fmt.Print("Выбор команды: ")
	_, err := fmt.Scan(&selectedUser)
	if err != nil {
		return "", fmt.Errorf("некорректный ввод")
	}

	if selectedUser == 0 {
		return "exit", nil
	}
	if selectedUser < 1 || selectedUser > len(keys) {
		return "", fmt.Errorf("некорректный выбор")
	}

	return keys[selectedUser-1], nil
}
