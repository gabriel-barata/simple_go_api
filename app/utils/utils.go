package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnvFile() error {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to load current working directory: %v", err)
		return err
	}

	envPath := filepath.Join(wd, "..", "..", ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Unable to load the .env file: %v", err)
		return err
	}

	return nil
}
