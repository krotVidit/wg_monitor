// Package ui
package ui

type CLI struct{}

func (CLI) SelectCommand(cmds map[string]string) (string, error) {
	return selectCommand(cmds)
}

func (CLI) Wrap(text, clr string) string {
	return wrap(text, clr)
}

func (CLI) BoldWrap(text string) string {
	return boldWrap(text)
}
