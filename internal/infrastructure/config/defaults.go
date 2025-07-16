package config

import "github.com/spf13/viper"

func (cfg *Config) setDefaults() {

	viper.SetDefault("APP.PORT", "8080")
	viper.SetDefault("APP.DEBUG", true)
	viper.SetDefault("APP.DBCONN", "")
}
