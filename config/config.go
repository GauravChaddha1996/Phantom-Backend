package config

import (
	"github.com/spf13/viper"
	"log"
)

type EnvConfig struct {
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

func ReadEnvConfig() *EnvConfig {
	// Make viper and read envConfig
	v := viper.New()
	v.AddConfigPath("./config/")
	v.SetConfigName("config")
	v.SetConfigType("json")
	viperConfigReadErr := v.ReadInConfig()
	if viperConfigReadErr != nil {
		log.Fatal(viperConfigReadErr)
	}

	// Read from viper into the envConfig struct
	var envConfig EnvConfig
	viperConfigUnMarshalErr := v.Unmarshal(&envConfig)
	if viperConfigUnMarshalErr != nil {
		log.Fatal(viperConfigUnMarshalErr)
	}
	return &envConfig
}
