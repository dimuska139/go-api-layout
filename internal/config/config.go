package config

import (
	"fmt"
	"github.com/dimuska139/urlshortener/internal/constants"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type DbConf struct {
	Dsn string `mapstructure:"dsn"`
}

type Config struct {
	Loglevel     string `mapstructure:"loglevel"`
	Env          string `mapstructure:"env"`
	JwtSecretKey string `mapstructure:"jwt_secret_key"`
	Port         int    `mapstructure:"port"`
	Domain       string `mapstructure:"domain"`
	Db           DbConf `mapstructure:"db"`
}

func NewConfig(configPath string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("reading config error: %v", err)
	}

	var cfg Config

	if err = yaml.Unmarshal(yamlFile, &cfg); err != nil {
		return nil, fmt.Errorf("parsing config error: %v", err)
	}

	if cfg.Port == 0 {
		cfg.Port = constants.DefaultPort
	}

	return &cfg, nil
}
