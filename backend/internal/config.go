package internal

import (
	"fmt"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	DBUser     string `env:"POSTGRES_USER" envDefault:"postgres"`
	DBPassword string `env:"POSTGRES_PASSWORD,required"`
	DBName     string `env:"POSTGRES_DB" envDefault:"postgres"`

	Port string `env:"APP_CONTAINER_PORT,required"`

	FrontendURL string `env:"FRONTEND_URL,required"`
}

func GetConfig() (*Config, error) {
	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("GetConfig: %w", err)
	}
	return cfg, nil
}
