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
	RABBIT_MQ_ADDR        string
	CREATE_DONATION_CH    string
	EDIT_DONATION_CH      string
	DELETE_DONATION_CH    string
	GOOGLE_MAPS_API_KEY   string
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
	CREATE_DONATION_CH = os.Getenv("CREATE_DONATION_CH")
	EDIT_DONATION_CH = os.Getenv("EDIT_DONATION_CH")
	DELETE_DONATION_CH = os.Getenv("DELETE_DONATION_CH")
	GOOGLE_MAPS_API_KEY = os.Getenv("GOOGLE_MAPS_API_KEY")
}
