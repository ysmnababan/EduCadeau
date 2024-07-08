package helper

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	RABBIT_MQ_ADDR        string
	USER_REGISTER_CHANNEL string
	USER_EDIT_CHANNEL     string
	PORT                  string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to get .env")
	}

	RABBIT_MQ_ADDR = os.Getenv("RABBIT_MQ_ADDR")
	USER_REGISTER_CHANNEL = os.Getenv("USER_REGISTER_CHANNEL")
	USER_EDIT_CHANNEL = os.Getenv("USER_EDIT_CHANNEL")
	PORT = os.Getenv("PORT")

}
