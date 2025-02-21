package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	StorageType string `yaml:"storage"`
	Address     string `yaml:"address"`
}

func New() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		// log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// log.Fatalf("config file does not exists: %s", configPath)
	}

	cfg := Config{}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
