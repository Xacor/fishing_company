package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Port      string `mapstructure:"PORT"`
	DBUrl     string `mapstructure:"DB_URL"`
	TestDBUrl string `mapstructure:"TEST_DB_URL"`
	Secret    string `mapstructure:"SECRET"`
}

func LoadConfig(path string) (c Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		err = fmt.Errorf("Config: %w", err)
	} else {
		log.Info("Trying to load settings from envs")
		viper.AutomaticEnv()
	}

	if err = viper.Unmarshal(&c); err != nil {
		err = fmt.Errorf("Unable to decode into struct, %v", err)
		return
	}
	return
}
