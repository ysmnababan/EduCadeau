package helper

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	USER_SERVICE_PORT string
	TOKEN_KEY         string
	USER_SERVICE_HOST string
	PORT              string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to get .env")
	}

	USER_SERVICE_PORT = os.Getenv("USER_SERVICE_PORT")
	TOKEN_KEY = os.Getenv("TOKEN_KEY")
	USER_SERVICE_HOST = os.Getenv("USER_SERVICE_HOST")
	PORT = os.Getenv("PORT")

}
