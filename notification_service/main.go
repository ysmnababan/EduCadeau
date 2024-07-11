package main

import (
	"log"
	"net/http"
	"notification_service/handler"
	helper "notification_service/helper"

	"github.com/labstack/echo"
)

func init() {
	// load the .env
	helper.LoadEnv()
	log.Println("env var loaded")
}
func main() {
	//setup the receiver
	go handler.NotifyUserRegister()
	go handler.NotifyUserEditData()
	go handler.NotifyCreateDonation()
	go handler.NotifyEditDonation()
	// http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte("OK"))
	// })

	// go func() {
	// 	log.Println("Starting HTTP server on port 8080")
	// 	if err := http.ListenAndServe(":"+helper.PORT, nil); err != nil {
	// 		log.Fatalf("Error starting HTTP server: %v", err)
	// 	}
	// }()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":" + helper.PORT))
	forever := make(chan bool)
	log.Println("NOTIFICATION SERVICE STARTED...")
	<-forever
}
