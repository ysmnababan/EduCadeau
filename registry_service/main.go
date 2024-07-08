package main

import (
	"context"
	"registry_service/config"
	"registry_service/controller"
	"registry_service/helper"
	"registry_service/repository"
	"registry_service/setup"
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
	registryController := &controller.RegistryController{
		RR:           repo,
		DonationGRPC: setup.SetupClientForDonationServer(),
		UserGRPC:     setup.SetupClientForUserServer(),
	}

	setup.SetupGPRCServer(registryController)

}
