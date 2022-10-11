package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Debug  bool `yaml:"debug"`
	Server struct {
		BindAddress string `yaml:"bind_address"`
		BindPort    string `yaml:"bind_port"`
	}
	Storage struct {
		DatabaseName string `yaml:"database_name"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{Debug: false}
	})
	return instance
}

func LoadData(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(content, GetConfig()); err != nil {
		return err
	}
	return nil
}
