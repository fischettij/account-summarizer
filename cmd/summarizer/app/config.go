package app

import (
	"os"
)

const (
	defaultPort = "8080"
)

type Config struct {
	Port           string
	FilesDirectory string
}

func LoadConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config := &Config{
		Port:           port,
		FilesDirectory: os.Getenv("FILES_DIRECTORY"),
	}

	return config, nil
}
