package env_vars

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DBHost    string
	DBPort    int
	DebugMode bool
}

func LoadConfig() (Config, error) {
	var config Config

	// Required: DB_HOST
	host, exists := os.LookupEnv("DB_HOST")
	if !exists || host == "" {
		return Config{}, fmt.Errorf("DB_HOST is required")
	}
	config.DBHost = host

	// Optional: DB_PORT
	config.DBPort = 5432

	if portStr, exists := os.LookupEnv("DB_PORT"); exists {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return Config{}, fmt.Errorf("invalid DB_PORT: %w", err)
		}
		config.DBPort = port
	}

	// Optional: DEBUG_MODE
	config.DebugMode = false

	if debugStr, exists := os.LookupEnv("DEBUG_MODE"); exists {
		debug, err := strconv.ParseBool(debugStr)
		if err != nil {
			return Config{}, fmt.Errorf("invalid DEBUG_MODE: %w", err)
		}
		config.DebugMode = debug
	}

	return config, nil
}
