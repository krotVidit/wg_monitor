// Package runner
package runner

import (
	"fmt"
	"sort"
)

// ANSI цвета для CLI
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
	colorBold   = "\033[1m"
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
			fmt.Println(colorRed, "Ошибка закрытия сессии:", e, colorReset)
		}
	}()

	fmt.Println(colorGreen + "✅ Подключено к серверу" + colorReset)

	cmds, err := r.commands.LoadCommand("commands.json")
	if err != nil {
		return "", fmt.Errorf("ошибка загрузки команд: %w", err)
	}

	for {
		key, err := selectCommand(cmds)
		if err != nil {
			fmt.Println(colorRed+"Ошибка:", err, colorReset)
			continue
		}

		if key == "exit" {
			fmt.Println(colorYellow + "👋 Выход из программы." + colorReset)
			break
		}

		fmt.Printf(colorCyan+"\n🚀 Выполняется команда: %s\n"+colorReset, key)
		fmt.Println(colorCyan + "====================================================" + colorReset)

		output, err := r.commands.RunCommand(client, cmds[key])
		if err != nil {
			fmt.Println(colorRed+"Ошибка выполнения команды:", err, colorReset)
			continue
		}

		fmt.Printf(colorGreen+"%s"+colorReset, output)
		fmt.Println(colorCyan + "====================================================" + colorReset)
	}

	return "Завершено.", nil
}

func selectCommand(cmds map[string]string) (string, error) {
	fmt.Println(colorBold + "\n========================================" + colorReset)
	fmt.Println(colorBold + "           Доступные команды" + colorReset)
	fmt.Println(colorBold + "========================================" + colorReset)

	keys := make([]string, 0, len(cmds))
	for name := range cmds {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	for i, name := range keys {
		fmt.Printf(colorCyan+"  %d."+colorReset+" %s\n", i+1, name)
	}
	fmt.Printf(colorYellow + "  0." + colorReset + " Выйти\n")

	var selectedUser int
	fmt.Print(colorBold + "\nВыбор команды: " + colorReset)
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
