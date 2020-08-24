package config

import "github.com/spf13/viper"

// InitConfig - parse config
func InitConfig() error {
	viper.AddConfigPath("../config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
