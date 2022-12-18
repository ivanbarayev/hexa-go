package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func exportConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	if os.Getenv("MODE") == "DOCKER" {
		viper.SetConfigName("config-docker.yml")
	} else {
		viper.SetConfigName("config.yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// ParseConfig Parse config file
func ParseConfig() (*Config, error) {
	if err := exportConfig(); err != nil {
		return nil, err
	}

	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
