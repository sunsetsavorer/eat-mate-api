package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort      string
	Debug        bool
	AllowOrigins []string
	DbConn       string
	JWTSecret    string
}

func NewConfig() *Config {

	return &Config{}
}

func (cfg *Config) LoadConfig() error {

	viper.SetConfigFile("config/.env")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config: %v", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %v", err)
	}

	return nil
}
