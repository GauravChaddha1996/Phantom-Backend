package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Database struct {
		Driver   string
		Username string
		Network  string
		Host     string
		Port     string
		Name     string
	}
	Cache struct {
		Network     string
		Host        string
		Port        int
		MaxIdle     int
		IdleTimeout int64
	}
}

func ReadConfig() *Config {
	// Make viper and read config
	v := viper.New()
	v.AddConfigPath("./config/")
	v.SetConfigName("config")
	v.SetConfigType("json")
	viperConfigReadErr := v.ReadInConfig()
	if viperConfigReadErr != nil {
		log.Fatal(viperConfigReadErr)
	}

	// Read from viper into the config struct
	var config Config
	viperConfigUnMarshalErr := v.Unmarshal(&config)
	if viperConfigUnMarshalErr != nil {
		log.Fatal(viperConfigUnMarshalErr)
	}
	return &config
}
