package main

import (
	"api_gateway/handler"
	"api_gateway/helper"
	"api_gateway/setup"
	"log"

	"github.com/labstack/echo/v4"
)

func init() {
	helper.LoadEnv()
}

func main() {
	e := echo.New()

	handler := &setup.Handler{
		User:     handler.InitUserHandler(),
		Donation: handler.InitDonationHandler(),
		Registry: handler.InitRegistryHandler(),
	}
	setup.SetupRESTServer(e, handler)

	log.Fatal(e.Start(":" + helper.PORT))
}
