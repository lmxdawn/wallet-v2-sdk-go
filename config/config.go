package config

import (
	"github.com/jinzhu/configor"
)

type AppConfig struct {
	Port uint `yaml:"port"`
}

type WalletConfig struct {
	Appid     string `yaml:"appid"`
	SecretKey string `yaml:"secret_key"`
	Url       string `yaml:"url"`
}

type Config struct {
	App    AppConfig
	Wallet WalletConfig
}

func NewConfig(confPath string) (Config, error) {
	var config Config
	if confPath != "" {
		err := configor.Load(&config, confPath)
		if err != nil {
			return config, err
		}
	} else {
		err := configor.Load(&config, "config/config-example.yml")
		if err != nil {
			return config, err
		}
	}
	return config, nil
}
