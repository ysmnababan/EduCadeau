package main

import (
	"log"
	"notification_service/handler"
	helper "notification_service/helper"
)

func main() {
	// load the .env
	helper.LoadEnv()

	//setup the receiver
	forever := make(chan bool)
	go handler.NotifyUserRegister()
	go handler.NotifyUserEditData()
	log.Println("NOTIFICATION SERVICE STARTED...")
	<-forever
}
