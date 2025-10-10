package command

import (
	"encoding/json"
	"os"
)

func LoadCommand(pathFile string) (map[string]string, error) {
	commands := make(map[string]string)
	configData, err := os.ReadFile(pathFile)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(configData, &commands); err != nil {
		return nil, err
	}
	return commands, nil
}
