package internal

import (
	"fmt"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	Port        string `env:"APP_CONTAINER_PORT,notEmpty"`
	DatabaseDSN string `env:"POSTGRES_DSN"`
	FrontendURL string `env:"FRONTEND_URL,notEmpty"`
}

func GetConfig() (*Config, error) {
	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("GetConfig: %w", err)
	}
	if len(cfg.DatabaseDSN) == 0 {
		dbEnv := struct {
			User     string `env:"POSTGRES_USER" envDefault:"postgres"`
			Password string `env:"POSTGRES_PASSWORD,required"`
			Name     string `env:"POSTGRES_DB" envDefault:"postgres"`
		}{}
		if err := env.Parse(&dbEnv); err != nil {
			return nil, fmt.Errorf("GetConfig: %w", err)
		}
		cfg.DatabaseDSN = fmt.Sprintf("postgres://%[1]v:%[2]v@db/%[3]v?sslmode=disable", dbEnv.User, dbEnv.Password, dbEnv.Name)
	}
	return cfg, nil
}
