package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func Loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("LOADED ENV FILE")
}
