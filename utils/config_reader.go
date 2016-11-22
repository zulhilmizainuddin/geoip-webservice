package utils

import (
	"github.com/spf13/viper"
)

const configName string = "config"
const configPath string = "./"

func ReadConfigFile() error {
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err == nil {
		viper.WatchConfig()
	}

	return err
}
