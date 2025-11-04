package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const configFileName = "config.yml"

type Config struct {
	Database struct {
		URL      string `yaml:"url"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

func getConfigFilePath() (string, error) {
	return "./" + configFileName, nil
	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	return "", err
	// }
	//
	// return filepath.Join(homeDir, configFileName), err
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
