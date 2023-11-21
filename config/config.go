package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Print("Error loading .env file")
		log.Panic()
	}

	// Return the value of the variable
	return os.Getenv(key)
}
