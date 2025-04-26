package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Debug        bool   `mapstructure:"DEBUG"`
	ServerPort   uint    `mapstructure:"SERVER_PORT"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`

	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     uint   `mapstructure:"POSTGRES_PORT"`
	PostgresSSLMode  string `mapstructure:"POSTGRES_SSL_MODE"`
	PostgresTimeZone string `mapstructure:"POSTGRES_TIME_ZONE"`
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
		viper.SetDefault("POSTGRES_USER", "admin")
		viper.SetDefault("POSTGRES_DB", "authify")
		viper.SetDefault("POSTGRES_HOST", "127.0.0.1")
		viper.SetDefault("POSTGRES_PORT", 5432)
		viper.SetDefault("POSTGRES_SSL_MODE", "disable")
		viper.SetDefault("POSTGRES_TIME_ZONE", "Asia/Tehran")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Failed to read config file: %v", err)
		}

		// check for requirement variables
		requiredVars := []string{"JWT_SECRET_KEY", "POSTGRES_PASSWORD"}
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
