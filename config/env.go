package config

import (
	"log"
	"os"
)

// Config holds the configuration for this process
type Config struct {
	APIKey      string
	Pin         string
	StoragePath string
}

// FromEnv reads the configuration from the environment variables
func FromEnv() *Config {
	return &Config{
		APIKey:      get("HK_API_KEY"),
		Pin:         get("HK_PIN"),
		StoragePath: get("HK_STORAGE_PATH"),
	}
}

func get(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalln("Missing environment variable:", k)
	}
	return v
}
