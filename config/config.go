package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config holds all application configuration
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string `mapstructure:"port"`
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

// Load loads configuration from YAML file
func Load() *Config {
	// Set config file name
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Allow environment variables to override
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode config: %v", err)
	}

	return &config
}
