package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	LOCALHOST         string
	USER              string
	PASSWORD          string
	DBNAME            string
	DB_PORT           string
	USER_SERVICE_PORT string
	TOKEN_KEY         string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to get .env")
	}

	LOCALHOST = os.Getenv("DB_HOST")
	USER = os.Getenv("DB_USER")
	PASSWORD = os.Getenv("DB_PASSWORD")
	DBNAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
	USER_SERVICE_PORT = os.Getenv("USER_SERVICE_PORT")
	TOKEN_KEY = os.Getenv("TOKEN_KEY")
	log.Println("localhost: ", LOCALHOST)
}
