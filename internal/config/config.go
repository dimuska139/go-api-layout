package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"

	"github.com/dimuska139/urlshortener/internal/service/server/rest"
	"github.com/dimuska139/urlshortener/internal/service/shrink"
	"github.com/dimuska139/urlshortener/internal/storage/postgresql/migrator"
	"github.com/dimuska139/urlshortener/pkg/postgresql"
)

type Config struct {
	Version string

	Loglevel   string            `yaml:"loglevel"`
	Env        string            `yaml:"env"`
	Shrink     shrink.Config     `yaml:"shrink"`
	PostgreSQL postgresql.Config `yaml:"postgresql"`
	Migrator   migrator.Config   `yaml:"migrator"`
	HttpServer rest.Config       `yaml:"http_server"`
}

type VersionParam string

func NewConfig(configPath string, version VersionParam) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	cfg.Version = string(version)

	return &cfg, nil
}
