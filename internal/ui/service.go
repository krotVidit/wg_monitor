// Package ui
package ui

type CLI struct{}

func (CLI) SelectCommand(cmds map[string]string) (string, error) {
	return selectCommand(cmds)
}
