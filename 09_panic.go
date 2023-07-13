package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string
	APIKey      string
}

func main() {
	// Load configuration
	config, err := loadConfig("config.yaml")
	if err != nil {
		panic(fmt.Errorf("Error: %w", err))
	}

	// Use the configuration
	fmt.Println("Database URL:", config.DatabaseURL)
	fmt.Println("API Key:", config.APIKey)
}

// loadConfig loads the configuration from the specified file using viper
func loadConfig(filePath string) (*Config, error) {
	// Initialize viper
	v := viper.New()
	v.SetConfigFile(filePath)

	// Read the configuration file
	err := v.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read configuration file: %w", err)
	}

	// Unmarshal the configuration into a struct
	var config Config
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	return &config, nil
}
