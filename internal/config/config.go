package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	r, err := os.ReadFile(configFilePath)

	if err != nil {
		return Config{}, err
	}

	var config Config

	if err := json.Unmarshal(r, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	return write(*c)
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil
}

func write(cfg Config) error {

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	if err := os.WriteFile(configFilePath, data, 0600); err != nil {
		return err
	}
	return nil
}
