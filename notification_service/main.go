package main

import (
	"fmt"
	"log"
	"net/http"
	"notification_service/handler"
	helper "notification_service/helper"
	"time"

	"github.com/labstack/echo"
	"github.com/robfig/cron/v3"
)

func init() {
	// load the .env
	helper.LoadEnv()
	log.Println("env var loaded")
}
func main() {
	//setup the receiver
	c := cron.New()

	// Schedule the job to run every 5 minutes
	c.AddFunc("* * * * *", func() {
		fmt.Println("Running job at:", time.Now())
		// Replace this with your actual job logic
		// Example: sendEmail()
		handler.SpamEmail("spam email", "INI SPAM")
	})

	c.Start()
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
