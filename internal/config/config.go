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
	DB struct {
		Postgres struct {
			Host     string `env:"POSTGRES_HOST"`
			Port     int    `env:"POSTGRES_PORT"`
			User     string `env:"POSTGRES_USER"`
			Password string `env:"POSTGRES_PASSWORD"`
			DBName   string `env:"POSTGRES_DB"`
			SSLMode  string `env:"POSTGRES_SSL_MODE"`
		}
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
