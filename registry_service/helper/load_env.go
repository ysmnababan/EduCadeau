package helper

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	USER_SERVICE_PORT     string
	DONATION_SERVICE_PORT string
	MONGO_URI             string
	USER_SERVICE_HOST     string
	DONATION_SERVICE_HOST string
	RABBIT_MQ_ADDR        string
	CREATE_REGISTRY_CH    string

)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to get .env")
	}

	DONATION_SERVICE_PORT = os.Getenv("DONATION_SERVICE_PORT")
	MONGO_URI = os.Getenv("MONGO_URI")
	USER_SERVICE_PORT = os.Getenv("USER_SERVICE_PORT")
	USER_SERVICE_HOST = os.Getenv("USER_SERVICE_HOST")
	RABBIT_MQ_ADDR = os.Getenv("RABBIT_MQ_ADDR")
	CREATE_REGISTRY_CH = os.Getenv("CREATE_REGISTRY_CH")
	DONATION_SERVICE_HOST = os.Getenv("DONATION_SERVICE_HOST")
}
