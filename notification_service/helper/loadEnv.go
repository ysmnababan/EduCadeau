package helper

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	RABBIT_MQ_ADDR   string
	PORT             string
	SENDGRID_API_KEY string

	USER_REGISTER_CHANNEL string
	USER_EDIT_CHANNEL     string

	CREATE_DONATION_CH string
	EDIT_DONATION_CH   string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to get .env")
	}

	RABBIT_MQ_ADDR = os.Getenv("RABBIT_MQ_ADDR")
	SENDGRID_API_KEY = os.Getenv("SENDGRID_API_KEY")

	PORT = os.Getenv("PORT")

	USER_REGISTER_CHANNEL = os.Getenv("USER_REGISTER_CHANNEL")
	USER_EDIT_CHANNEL = os.Getenv("USER_EDIT_CHANNEL")

	CREATE_DONATION_CH = os.Getenv("CREATE_DONATION_CH")
	EDIT_DONATION_CH = os.Getenv("EDIT_DONATION_CH")

}
