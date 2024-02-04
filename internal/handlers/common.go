package handlers

import (
	"path/filepath"
	"regexp"
	"strings"
)

func isCSVFilePath(filePathOrName string) bool {
	ext := filepath.Ext(filePathOrName)
	ext = strings.ToLower(ext)

	return ext == ".csv"
}

func isValidEmail(email string) bool {
	// This is a simple regex. Will be replaced for one more complex.
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return emailRegex.MatchString(email)
}
