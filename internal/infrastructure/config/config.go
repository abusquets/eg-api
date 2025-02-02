package config

import (
	"fmt"
	"os"
	"sync"
)

// AppConfig struct stores the application configuration
type AppConfig struct {
	ServerPort              string
	ServerUri               string
	MongoDBUri              string
	MongoDB                 string
	RedisUrl                string
	RedisPassword           string
	RedisDB                 int
	AccessTokenDuration     int
	RefresTokenDuration     int
	ApiName                 string
	ApiVersion              string
	CookieSSL               bool
	InboundEventsStreamName string
}

var (
	config     *AppConfig // Global variable to hold the singleton instance
	configOnce sync.Once  // Ensures config is initialized only once
)

// getEnv retrieves environment variables or returns a default value if not set
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetAppConfig initializes the configuration (singleton)
func GetAppConfig() (*AppConfig, error) {
	fmt.Println("Loading configuration")

	var err error

	// Ensure the configuration is created only once
	configOnce.Do(func() {
		config = &AppConfig{
			ServerPort: getEnv("SERVER_PORT", "8080"),
			ApiName:    getEnv("API_NAME", "My API"),
			ApiVersion: getEnv("API_VERSION", "1.0.0"),
		}

	})

	// Return the singleton instance and any error
	return config, err
}
