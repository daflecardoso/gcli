package config

import (
	"encoding/json"
	"os"
)

const FileName = "gcli.json"

type Config struct {
	Name         string   `json:"name"`
	Color        string   `json:"color"`
	ShowTutorial bool     `json:"showTutorial"`
	Scopes       []string `json:"scopes"`
}

func Load() (*Config, error) {
	data, err := os.ReadFile(FileName)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
