package main

import (

	"notification_service/config"
	"notification_service/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	// load the .env
	helper.LoadEnv()
	//setup the receiver
	forever := make(chan bool)
	go handler.NotifyUserRegister()
	go handler.NotifyUserEditData()
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	go func() {
		log.Println("Starting HTTP server on port 8080")
		if err := http.ListenAndServe(":"+helper.PORT, nil); err != nil {
			log.Fatalf("Error starting HTTP server: %v", err)
		}
	}()

	log.Println("NOTIFICATION SERVICE STARTED...")
	<-forever
}
