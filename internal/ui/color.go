// Package ui
package ui

const (
	Reset = "\033[0m"
	Bold  = "\033[1m"

	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
)

func wrap(text, clr string) string {
	mapping := map[string]string{
		"red":    Red,
		"green":  Green,
		"yellow": Yellow,
		"cyan":   Cyan,
		"bold":   Bold,
	}[clr]

	if mapping == "" {
		mapping = Reset
	}

	return mapping + text + Reset
}

func boldWrap(text string) string {
	return Bold + text + Reset
}
