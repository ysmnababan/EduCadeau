package helper

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	USER_SERVICE_PORT     string
	DONATION_SERVICE_PORT string
	MONGO_URI             string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to get .env")
	}

	DONATION_SERVICE_PORT = os.Getenv("DONATION_SERVICE_PORT")
	MONGO_URI = os.Getenv("MONGO_URI")
	USER_SERVICE_PORT = os.Getenv("USER_SERVICE_PORT")
}
