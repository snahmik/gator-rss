package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (*Config, error) {
	var config Config

	filePath, err := getConfigFilePath()
	if err != nil {
		return &config, err
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return &config, fmt.Errorf("error reading config file: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(file))
	if err := decoder.Decode(&config); err != nil {
		return &config, fmt.Errorf("error parsing config file: %w", err)
	}

	return &config, nil
}

func (c *Config) SetUser(username string) (*Config, error) {
	c.CurrentUserName = username

	filePath, err := getConfigFilePath()
	if err != nil {
		return c, err
	}

	configBytes, err := json.Marshal(c)
	if err != nil {
		return c, fmt.Errorf("error marshalling config struct: %w", err)
	}

	err = os.WriteFile(filePath, configBytes, 0644)
	if err != nil {
		return c, fmt.Errorf("error writing config file: %w", err)
	}

	return c, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error reading home directory: %w", err)
	}
	configPath := filepath.Join(homeDir, ".gatorconfig.json")

	return configPath, nil
}
