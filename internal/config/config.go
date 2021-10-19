package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/strpc/resume-success/pkg/logging"
)

type Config struct {
	App struct {
		Port     int             `env:"PORT"`
		LogLevel string          `env:"LOG_LEVEL"`
		LogType  logging.LogType `env:"LOG_TYPE"`
	}
}

func GetConfig() *Config {
	var once sync.Once
	config := &Config{}
	once.Do(func() {
		if err := cleanenv.ReadEnv(config); err != nil {
			log.Fatal(err)
		}
	})
	return config
}
