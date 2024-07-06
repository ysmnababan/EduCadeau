package controller

import (
	"context"
	"donation_service/pb"
	"donation_service/repository"
)

type DonationController struct {
	DC repository.DonationRepo
}

func (c *DonationController) GetAllDonations(ctx context.Context, in *pb.DonationReq) (*pb.DonationList, error) {

	return nil, nil
}

func (c *DonationController) GetDonationDetail(ctx context.Context, in *pb.DonationDetailReq) (*pb.DonationDetailResp, error) {

	return nil, nil
}

func (c *DonationController) CreateDonation(ctx context.Context, in *pb.CreateDonationReq) (*pb.CreateResp, error) {

	return nil, nil
}

func (c *DonationController) EditDonation(ctx context.Context, in *pb.EditDonationReq) (*pb.EditResp, error) {

	return nil, nil
}

func (c *DonationController) DeleteDonation(ctx context.Context, in *pb.DeleteDonationReq) (*pb.DeleteResp, error) {

	return nil, nil
}
