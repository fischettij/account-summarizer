package app

import (
	"github.com/fischettij/account-summarizer/internal/email"
	"os"
)

const (
	defaultPort = "8080"
)

type Config struct {
	Port           string
	FilesDirectory string
	SMTP           email.Config
}

func LoadConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	smtp := email.Config{
		Port:        os.Getenv("SMTP_PORT"),
		ServerURL:   os.Getenv("SMTP_SERVER_URL"),
		From:        os.Getenv("SMTP_FROM"),
		Username:    os.Getenv("SMTP_USERNAME"),
		Password:    os.Getenv("SMTP_PASSWORD"),
		Identity:    os.Getenv("SMTP_IDENTITY"),
		TLSHostName: os.Getenv("SMTP_TLS_HOSTNAME"),
	}

	config := &Config{
		Port:           port,
		FilesDirectory: os.Getenv("FILES_DIRECTORY"),
		SMTP:           smtp,
	}

	return config, nil
}
