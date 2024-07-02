package controller

import (
	"context"
	"log"
	"user_service/helper"
	"user_service/models"
	"user_service/pb"
)

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
	resp, err := s.UserRepo.GetAllUser()
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	out := []*pb.UserDetail{}
	for _, val := range resp {
		user := &pb.UserDetail{}
		user.UserId = uint64(val.UserID)
		user.Username = val.Username
		user.Email = val.Email
		user.Deposit = val.Deposit
		user.Fname = val.Fname
		user.Lname = val.Lname
		user.Address = val.Address
		user.Age = int64(val.Age)
		user.PhoneNumber = val.PhoneNumber
		user.ProfilePictureUrl = val.ProfilePictureUrl
		out = append(out, user)
	}

	log.Printf("GET ALL USER SUCCESS\n====================\n\n\n")

	return &pb.GetAllResp{
		List: out,
	}, nil
}

func (s *UserController) GetUserDetail(ctx context.Context, in *pb.DetailReq) (*pb.UserDetailResp, error) {
	val, err := s.UserRepo.GetInfo(uint(in.UserId))
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}
	user := &pb.UserDetailResp{}
	user.UserId = uint64(val.UserID)
	user.Username = val.Username
	user.Email = val.Email
	user.Deposit = val.Deposit
	user.Fname = val.Fname
	user.Lname = val.Lname
	user.Address = val.Address
	user.Age = int64(val.Age)
	user.PhoneNumber = val.PhoneNumber
	user.ProfilePictureUrl = val.ProfilePictureUrl

	log.Printf("GET USER DETAIL SUCCESS\n====================\n\n\n")
	return user, nil
}
func (s *UserController) TopUp(ctx context.Context, in *pb.TopUpReq) (*pb.TopUpResp, error) {
	if in.Amount <= 0 || in.Amount >= 20000 {
		return nil, helper.ParseErrorGRPC(helper.ErrParam)
	}
	respU, err := s.UserRepo.TopUp(uint(in.UserId), in.Amount)
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	log.Println("new balance: ", respU)
	log.Printf("TOP UP SUCCESS\n====================\n\n\n")
	return &pb.TopUpResp{Balance: respU}, nil

}
func (s *UserController) EditDataUser(ctx context.Context, in *pb.EditReq) (*pb.EditResp, error) {
	var user models.UserUpdateRequest

	user.Username = in.Username
	user.Fname = in.Fname
	user.Lname = in.Lname
	user.Address = in.Address
	user.Age = int(in.Age)
	user.PhoneNumber = in.PhoneNumber
	user.ProfilePictureUrl = in.ProfilePictureUrl

	respU, err := s.UserRepo.Update(uint(in.UserId), user)
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	out := &pb.EditResp{}
	out.UserId = uint64(respU.UserID)
	out.Username = respU.Username
	out.Email = respU.Email
	out.Deposit = respU.Deposit
	out.Fname = respU.Fname
	out.Lname = respU.Lname
	out.Address = respU.Address
	out.Age = int64(respU.Age)
	out.PhoneNumber = respU.PhoneNumber
	out.ProfilePictureUrl = respU.ProfilePictureUrl

	log.Printf("UPDATE USER SUCCESS\n====================\n\n\n")

	return out, nil
}
