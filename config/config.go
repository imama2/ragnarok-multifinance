package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MySQL struct {
		DSN string
	}
	Redis struct {
		Addr     string
		Password string
		DB       int
	}
	JWT struct {
		Secret string
	}
	Server struct {
		Port string
	}
}

func Load() (*Config, error) {
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
