package config

import (
	"errors"

	"github.com/codingconcepts/env"
	"go.uber.org/zap"
)

type Config struct {
	Addr string `env:"ADDR" require:"true"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Set(cfg); err != nil {
		zap.S().Errorf("new config from env failed, err: %v", err)
		return nil, err
	}
	if cfg.Addr == "" {
		zap.S().Errorf("env ADDR is empty")
		return nil, errors.New("env ADDR is empty")
	}
	return cfg, nil
}
