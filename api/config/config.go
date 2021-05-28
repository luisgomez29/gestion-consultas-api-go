package config

import (
	"log"

	"github.com/joho/godotenv"
)

// Load lee las configuraciones del archivo .env
func Load() (env map[string]string) {
	env, err := godotenv.Read()

	if err != nil {
		log.Fatal("Error reading .env file")
	}

	return
}
