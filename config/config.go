package config

import (
	"os"
	"strconv"
)

type Config struct {
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	ServerPort     string
	DBMaxOpenConns int
	DBMaxIdleConns int
	JWTSecret      string
	Environment    string
}

func Load() Config {
	return Config{
		DBHost:         getEnvOrDefault("MYSQL_HOST", "localhost"),
		DBPort:         getEnvOrDefault("MYSQL_PORT", "3306"),
		DBUser:         os.Getenv("MYSQL_USER"),
		DBPassword:     os.Getenv("MYSQL_PASSWORD"),
		DBName:         os.Getenv("MYSQL_DATABASE"),
		ServerPort:     getEnvOrDefault("SERVER_PORT", "8085"),
		DBMaxOpenConns: getEnvOrDefaultInt("DB_MAX_OPEN_CONNS", 10),
		DBMaxIdleConns: getEnvOrDefaultInt("DB_MAX_IDLE_CONNS", 5),
		JWTSecret:      getEnvOrDefault("JWT_SECRET", "your-secret-key"),
		Environment:    getEnvOrDefault("ENVIRONMENT", "development"),
	}
}

func getEnvOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvOrDefaultInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}
