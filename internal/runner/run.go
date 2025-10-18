// Package runner
package runner

import (
	"fmt"
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
	ui        UI
}

func New(connector SSHConnector, commands CommandRunner, ui UI) *Runner {
	return &Runner{
		connector: connector,
		commands:  commands,
		ui:        ui,
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
		key, err := r.ui.SelectCommand(cmds)
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
