// Package domain описывает структуру из JSON файла
package domain

import (
	"encoding/json"
	"os"
)

type Config struct {
	User    string `json:"user"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	PathKey string `json:"private_key_path"`
}

func LoadConfig(pathFile string) (*Config, error) {
	var cfg Config
	configData, err := os.ReadFile(pathFile)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(configData, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
