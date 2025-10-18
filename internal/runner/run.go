// Package runner
package runner

import (
	"fmt"
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
			fmt.Println(r.ui.Wrap(fmt.Sprintf("Ошибка закрытия сессии: %v", e), "red"))
		}
	}()

	fmt.Println(r.ui.Wrap("✅ Подключено к серверу", "green"))

	cmds, err := r.commands.LoadCommand("commands.json")
	if err != nil {
		return "", fmt.Errorf("ошибка загрузки команд: %w", err)
	}

	for {
		key, err := r.ui.SelectCommand(cmds)
		if err != nil {
			fmt.Println(r.ui.Wrap(fmt.Sprintf("Ошибка: %v", err), "red"))
			continue
		}

		if key == "exit" {
			fmt.Println(r.ui.Wrap("👋 Выход из программы.", "yellow"))
			break
		}

		fmt.Println(r.ui.Wrap(fmt.Sprintf("\n🚀 Выполняется команда: %s", key), "cyan"))
		fmt.Println(r.ui.Wrap("====================================================", "cyan"))

		output, err := r.commands.RunCommand(client, cmds[key])
		if err != nil {
			fmt.Println(r.ui.Wrap(fmt.Sprintf("Ошибка выполнения команды: %v", err), "red"))
			continue
		}

		fmt.Print(r.ui.Wrap(output, "green"))
		fmt.Println(r.ui.Wrap("====================================================", "cyan"))
	}

	return "Завершено.", nil
}
