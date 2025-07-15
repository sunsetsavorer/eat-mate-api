package config

import "github.com/spf13/viper"

func (cfg *Config) setDefaults() {

	// Here we set config variables default values
	viper.SetDefault("AllowOrigins", []string{"http://localhost:3000"})
}
