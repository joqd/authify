package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Debug        bool   `mapstructure:"DEBUG"`
	ServerPort   int    `mapstructure:"SERVER_PORT"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

var (
	configInstance *Config
	once           sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()

		viper.SetDefault("DEBUG", false)
		viper.SetDefault("SERVER_PORT", 8080)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Failed to read config file: %v", err)
		}

		// check for requirement variables
		requiredVars := []string{"JWT_SECRET_KEY"}
		for _, key := range requiredVars {
			if !viper.IsSet(key) {
				log.Fatalf("Missing required configuration: %s", key)
			}
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			log.Fatalf("Failed to unmarshal config file: %v", err)
		}
	})

	return configInstance
}
