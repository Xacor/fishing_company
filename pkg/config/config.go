package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBUrl  string `mapstructure:"DB_URL"`
	Secret string `mapstructure:"SECRET"`
}

func LoadConfig(path string) (c Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		err = fmt.Errorf("Config: %w", err)
		return
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		err = fmt.Errorf("Unable to decode into struct, %v", err)
		return
	}
	return
}
