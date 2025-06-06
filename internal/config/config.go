package config

import (
	"encoding/json"
	"os"
	"os/user"
)

const configFileName = ".gatorconfig.json"

func Read() Config {
	configFile, err := getConfigFilePath()
	if err != nil {
		return Config{}
	}
	jsonData, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}
	}
	newConfig := Config{}
	err = json.Unmarshal(jsonData, &newConfig)
	if err != nil {
		return Config{}
	}
	return newConfig
}

func (c *Config) SetUser(username string) error {
	if username == "" {
		currentUser, err := user.Current()
		if err != nil {
			return err
		}
		username = currentUser.Username
	}
	c.CurrentUserName = username
	err := write(*c)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homeDir + "/" + configFileName, nil
}

func write(cfg Config) error {
	newCfg, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	configFile, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(configFile, newCfg, 0600)
	if err != nil {
		return err
	}

	return nil
}
