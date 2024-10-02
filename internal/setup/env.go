package setup

import (
	"log"

	"github.com/joho/godotenv"
)

func MustEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("failed to load .env file: %v", err)
	}
}
