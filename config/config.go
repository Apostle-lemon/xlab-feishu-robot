package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var C Config

func ReadConfig() {
	viper.SetConfigName("config") // set the config file name. Viper will automatically detect the file extension name
	viper.AddConfigPath("./")     // search the config file under the current directory

	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		logrus.Error("Failed to unmarshal config")
	}

	logrus.Info("Configuration file loaded")
}
