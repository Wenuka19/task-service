package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port string
}

var AppConfig *Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, falling back to system env vars")
	}

	AppConfig = &Config{
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}
