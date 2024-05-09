package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RedisConfig struct {
		Addr     string `yaml:"addr"` //host:port Addr
		Password string `yaml:"password"`
	} `yaml:"redis"`
}

func ReadConfig() (*Config, error) {
	config := &Config{}
	config_bytes, err := os.ReadFile("config.yaml")

	if err != nil {
		return nil, errors.New("can't read the config file")
	}

	if err := yaml.Unmarshal(config_bytes, config); err != nil {
		return nil, errors.New("can't marshal config file")
	}

	return config, nil

}
