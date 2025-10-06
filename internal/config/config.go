package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	Port            string
	MongoDBURI      string
	MongoDBDatabase string
	JWTSecret       string
	APIKey          string
	LogLevel        string
	Environment     string
}

// LoadConfig loads the application configuration
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	return &Config{
		Port:            getEnv("PORT", "8080"),
		MongoDBURI:      getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		MongoDBDatabase: getEnv("MONGODB_DATABASE", "sailing_db"),
		JWTSecret:       getEnv("JWT_SECRET", "default-secret-change-this"),
		APIKey:          getEnv("API_KEY", "default-api-key-change-this"),
		LogLevel:        getEnv("LOG_LEVEL", "info"),
		Environment:     getEnv("ENV", "development"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
