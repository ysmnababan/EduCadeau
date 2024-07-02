package controller

import (
	"context"
	"log"
	"user_service/helper"
	"user_service/models"
	"user_service/pb"
	"user_service/repository"
)

type UserController struct {
	repository.UserRepo
}

func (s *UserController) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	var GetU models.User

	//validate user
	if in.Email == "" || in.Password == "" || in.Username == "" || in.Role == "" {
		return nil, helper.ParseErrorGRPC(helper.ErrParam)
	}

	// validate role
	if in.Role != "recipient" && in.Role != "admin" && in.Role != "donor" {
		return nil, helper.ParseErrorGRPC(helper.ErrParam)
	}

	GetU.Email = in.Email
	GetU.Password = in.Password
	GetU.Username = in.Username
	GetU.Role = in.Role

	respU, err := s.UserRepo.Register(GetU)
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	log.Printf("REGISTER SUCCESS\n====================\n\n\n")

	return &pb.RegisterResp{
		UserId:   uint64(respU.UserID),
		Username: respU.Username,
		Email:    respU.Email,
		Deposit:  respU.Deposit,
	}, nil

}

func (s *UserController) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	var GetU models.User

	//validate user
	if in.Email == "" || in.Password == "" {
		return nil, helper.ParseErrorGRPC(helper.ErrParam)
	}

	GetU.Email = in.Email
	GetU.Password = in.Password

	tokenString, err := s.UserRepo.Login(GetU)
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	log.Printf("LOGIN SUCCESS\n====================\n\n\n")

	return &pb.LoginResp{Token: tokenString}, nil
}
func (s *UserController) GetAllUser(ctx context.Context, in *pb.Req) (*pb.GetAllResp, error) {
	return nil, nil
}
func (s *UserController) GetUserDetail(ctx context.Context, in *pb.DetailReq) (*pb.UserDetailResp, error) {
	return nil, nil
}
func (s *UserController) TopUp(ctx context.Context, in *pb.TopUpReq) (*pb.TopUpResp, error) {
	return nil, nil
}
func (s *UserController) EditDataUser(ctx context.Context, in *pb.EditReq) (*pb.EditResp, error) {
	return nil, nil
}
