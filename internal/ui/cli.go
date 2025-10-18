// Package ui
package ui

import (
	"fmt"
	"sort"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
	colorBold   = "\033[1m"
)

func selectCommand(cmds map[string]string) (string, error) {
	printHeader()
	keys := sortKeys(cmds)
	printCommandList(keys)

	selectedUser, err := readUserSelection(len(keys))
	if err != nil {
		return "", err
	}

	if selectedUser == 0 {
		return "exit", nil
	}

	return keys[selectedUser-1], nil
}

func printHeader() {
	fmt.Println(colorBold + "\n========================================" + colorReset)
	fmt.Println(colorBold + "           Доступные команды" + colorReset)
	fmt.Println(colorBold + "========================================" + colorReset)
}

func sortKeys(cmds map[string]string) []string {
	keys := make([]string, 0, len(cmds))
	for name := range cmds {
		keys = append(keys, name)
	}
	sort.Strings(keys)
	return keys
}

func printCommandList(keys []string) {
	for i, name := range keys {
		fmt.Printf(colorCyan+"  %d."+colorReset+" %s\n", i+1, name)
	}
	fmt.Printf(colorYellow + "  0." + colorReset + " Выйти\n")
}

func readUserSelection(max int) (int, error) {
	var selected int
	fmt.Print(colorBold + "\nВыбор команды: " + colorReset)
	_, err := fmt.Scan(&selected)
	if err != nil {
		return 0, fmt.Errorf("ввод команды: %w", err)
	}
	if selected < 0 || selected > max {
		return 0, fmt.Errorf("некорректный выбор: %d", selected)
	}
	return selected, nil
}
