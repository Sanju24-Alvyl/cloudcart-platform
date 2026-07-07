package config

import "os"

type Config struct {
	AppName     string
	Environment string
	Port        string
	LogLevel    string
}

func Load() *Config {
	return &Config{
		AppName:     getEnv("APP_NAME", "api-gateway"),
		Environment: getEnv("ENVIRONMENT", "development"),
		Port:        getEnv("PORT", "8080"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}