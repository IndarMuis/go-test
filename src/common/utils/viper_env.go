package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func ViperEnv(key string) string {

	var config *viper.Viper = viper.New()

	config.SetConfigFile(".env")
	config.AutomaticEnv()

	// Find and read the config file
	err := config.ReadInConfig()

	if err != nil {
		fmt.Println("ERROR 1")
		return os.Getenv(key)
	}

	value, ok := config.Get(key).(string)

	if !ok {
		fmt.Println("ERROR 1")
		return os.Getenv(key)
	}

	return value
}
