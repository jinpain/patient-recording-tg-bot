package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Token      string      `yaml:"token"`
	Registrars []registrar `yaml:"registrars"`
}

type registrar struct {
	ChatId string `yaml:"chat_id"`
}

func MustLoad() *Config {
	configPath := "config/config.yaml"
	// configPath := os.Getenv("CONFIG_PATH")
	// if configPath == "" {
	// 	panic("config path is empty")
	// }

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
