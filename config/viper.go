package config

import "github.com/spf13/viper"

func NewViper() *viper.Viper {
	config := viper.New()
	config.SetConfigType("env")
	config.AddConfigPath(".")
	config.SetConfigName(".env")
	config.AutomaticEnv()
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return config
}
