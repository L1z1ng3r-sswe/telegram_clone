package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env             string        `yaml:"env" env-default:"local"`
	PostgresPath    string        `yaml:"postgres_path" env-required:"true"`
	GRPC            GRPC          `yaml:"grpc"`
	REST            REST          `yaml:"rest"`
	AccessTokenExp  time.Duration `yaml:"access_token_exp" env-required:"true"`
	RefreshTokenExp time.Duration `yaml:"refresh_token_exp" env-required:"true"`
	SecretKey       string        `yaml:"secret_key" env-required:"true"`
}

type GRPC struct {
	Port    int           `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

type REST struct {
	Port string `yaml:"port" env-required:"true"`
}

func MustLoad() *Config {
	configPath := "./config/local.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("Config-file is not exist: " + err.Error())
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("Failed to read config: " + err.Error())
	}

	return &cfg
}
