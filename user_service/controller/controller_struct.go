package controller

import "user_service/repository"

type UserController struct {
	repository.UserRepo
}

type UserDonation struct {
	UD repository.UserRepo
}

type UserRegistry struct {
	UR repository.UserRepo
}