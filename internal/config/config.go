package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DatabaseURL             string
	JWTSecret               string
	JWTExpirationHours      int
	FrontendURL             string
	ResetTokenExpirationMin int
	Port                    string
}

func Load() (*Config, error) {
	jwtExpHours, err := strconv.Atoi(getEnv("JWT_EXPIRATION_HOURS", "24"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_EXPIRATION_HOURS: %v", err)
	}

	resetTokenExpMin, err := strconv.Atoi(getEnv("RESET_TOKEN_EXPIRATION_MINUTES", "30"))
	if err != nil {
		return nil, fmt.Errorf("invalid RESET_TOKEN_EXPIRATION_MINUTES: %v", err)
	}

	return &Config{
		DatabaseURL:             getEnv("DATABASE_URL", ""),
		JWTSecret:               getEnv("JWT_SECRET", ""),
		JWTExpirationHours:      jwtExpHours,
		FrontendURL:             getEnv("FRONTEND_URL", "http://localhost:5173"),
		ResetTokenExpirationMin: resetTokenExpMin,
		Port:                    getEnv("PORT", "8080"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
