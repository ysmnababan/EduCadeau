package main

import (
	"user_service/config"
	"user_service/controller"
	"user_service/helper"
	"user_service/repository"
	"user_service/setup"
)

func init() {
	helper.LoadEnv()
}

func main() {
	// connect to db
	db := config.Connect()

	repo := &repository.Repo{DB: db}
	userController := &controller.UserController{
		UserRepo: repo,
	}

	setup.SetupGPRCServer(userController)

}
