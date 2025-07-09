package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Read() *Config {
	v := viper.New()

	v.AddConfigPath("./config")
	v.SetConfigName("config")
	v.SetConfigType("json")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config := new(Config)
	if err := v.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unmarshalling config error: %w", err))
	}

	return config
}
