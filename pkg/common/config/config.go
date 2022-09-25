package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port  string `mapstructure:"PORT"`
	DBUrl string `mapstructure:"DB_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/common/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		err = fmt.Errorf("fatal error config file: %w", err)
		return Config{}, err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		err = fmt.Errorf("unable to decode into struct, %v", err)
		return Config{}, err
	}
	return
}
