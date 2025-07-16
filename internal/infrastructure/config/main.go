package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Port         string
		Debug        bool
		AllowOrigins []string
		DbConn       string
	}
	JWT struct {
		Secret            string
		LifetimeInMinutes time.Duration
	}
}

func NewConfig() *Config {

	return &Config{}
}

func (cfg *Config) LoadConfig() error {

	cfg.setDefaults()

	viper.SetConfigFile("config/config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config: %v", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %v", err)
	}

	return nil
}
