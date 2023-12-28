package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		Nats `yaml:"nats"`
		PG   `yaml:"pg"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		Build   string `env-required:"true" yaml:"build"   env:"APP_BUILD"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	Nats struct {
		URL   string `env-required:"true" yaml:"NATS_URL" env:"NATS_URL"`
		Theme string `env-required:"true" yaml:"NATS_THEME" env:"NATS_THEME"`
		Timer string `env-required:"true" yaml:"NATS_TIMER" env:"NATS_TIMER"`
	}

	PG struct {
		URL string `env-required:"true" yaml:"PG_URL" env:"PG_URL"`
	}

	//SqlDB struct {
	//	URL string `env-required:"true" yaml:"SQL_URL" env:"SQL_URL"`
	//}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
