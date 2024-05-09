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

type GlobalConfig struct {
	NPubs int `yaml:"n_pubs"`
	NSubs int `yaml:"n_subs"`
	NMsgs int `yaml:"n_messages"`
}

func ReadConfig() (*Config, error) {
	config := &Config{}
	config_bytes, err := os.ReadFile("config.yaml")

	if err != nil {
		return nil, errors.New("can't read the config file")
	}

	if err := yaml.Unmarshal(config_bytes, config); err != nil {
		return nil, errors.New("can't unmarshal config file")
	}

	return config, nil

}

func ReadGlobalConfig(file_path string) (*GlobalConfig, error) {
	config := &GlobalConfig{}
	config_bytes, err := os.ReadFile(file_path)
	if err != nil {
		return nil, errors.New("can't load the global config file")
	}

	if err := yaml.Unmarshal(config_bytes, config); err != nil {
		return nil, errors.New("can't unmarshal global config file")
	}

	return config, nil
}
