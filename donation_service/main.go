package main

import (
	"context"
	"donation_service/config"
	"donation_service/controller"
	"donation_service/helper"
	"donation_service/repository"
	"donation_service/setup"
)

func init() {
	helper.LoadEnv()
}

func main() {
	// connect to db
	client, db := config.Connect(context.TODO(), "edu_cadeau")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	repo := &repository.Repo{DB: db}
	donationController := &controller.DonationController{
		DC:       repo,
		UserGRPC: setup.SetupClientForUserServer(),
	}

	setup.SetupGPRCServer(donationController)

}
