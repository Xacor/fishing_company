package config

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	Port    string `mapstructure:"PORT"`
	DBUrl   string `mapstructure:"DB_URL"`
	LogO    string `mapstructure:"LOGGER_TYPE"` //posible values: file, console, all
	LogFile string `mapstructure:"LOG_FILE"`    //filename with extension
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./envs")
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

func CustomLogFormatter(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] %s %s %s %d %s \"%s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.ErrorMessage,
	)
}
