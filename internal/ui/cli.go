// Package ui
package ui

import (
	"fmt"
	"sort"
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
	fmt.Println(boldWrap("\n========================================"))
	fmt.Println(boldWrap("           Доступные команды"))
	fmt.Println(boldWrap("========================================"))
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
		fmt.Printf("%s %s\n", wrap(fmt.Sprintf("  %d.", i+1), "cyan"), name)
	}
	fmt.Println(wrap("  0.", "yellow") + " Выйти")
}

func readUserSelection(max int) (int, error) {
	var selected int
	fmt.Print(boldWrap("\nВыбор команды: "))
	_, err := fmt.Scan(&selected)
	if err != nil {
		return 0, fmt.Errorf("ввод команды: %w", err)
	}
	if selected < 0 || selected > max {
		return 0, fmt.Errorf("некорректный выбор: %d", selected)
	}
	return selected, nil
}
