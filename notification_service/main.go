package main

import (
	"notification_service/config"
	"notification_service/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	config.Loadenv()
}

func main() {
	// Initialize Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/send-email", handler.SendEmail)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
