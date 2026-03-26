package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Values  Values
	DB      DB
	Logging Logging
}

func Load() (*Config, error) {
	var cfg Config

	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		return nil, err
	}

	if err := cfg.Values.validate(); err != nil {
		return nil, err
	}
	if err := cfg.DB.validate(); err != nil {
		return nil, err
	}
	if err := cfg.Logging.validate(); err != nil {
		return nil, err
	}

	cfg.Values.HttpAddr = ":8080"

	if err := cfg.Logging.parse(cfg.Values.AppEnv); err != nil {
		return nil, err
	}

	return &cfg, nil
}
