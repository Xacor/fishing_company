package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	DBUrl      string `mapstructure:"DB_URL"`
	TestDBUrl  string `mapstructure:"TEST_DB_URL"`
	Secret     string `mapstructure:"SECRET"`
	LoggingURL string `mapstructure:"LOGGING_URL"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	var c Config

	if err := viper.ReadInConfig(); err != nil {
		log.Error(err)
		return Config{}, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		log.Error(err)
		return Config{}, err
	}
	return c, nil
}
