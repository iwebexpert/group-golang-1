package config

import "github.com/spf13/viper"

//Parse - parse config
func Parse() error {
	viper.AddConfigPath("../config")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
