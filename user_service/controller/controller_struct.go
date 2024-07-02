package controller

import "user_service/repository"

type UserController struct {
	repository.UserRepo
}
