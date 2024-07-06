package controller

import (
	"context"
	"donation_service/pb/donation_rest"
	"donation_service/pb/user_donation"
	"donation_service/repository"
)

type DonationController struct {
	DC       repository.DonationRepo
	UserGRPC user_donation.UserDonationClient
}

func (c *DonationController) GetAllDonations(ctx context.Context, in *donation_rest.DonationReq) (*donation_rest.DonationList, error) {

	return nil, nil
}

func (c *DonationController) GetDonationDetail(ctx context.Context, in *donation_rest.DonationDetailReq) (*donation_rest.DonationDetailResp, error) {

	return nil, nil
}

func (c *DonationController) CreateDonation(ctx context.Context, in *donation_rest.CreateDonationReq) (*donation_rest.CreateResp, error) {

	return nil, nil
}

func (c *DonationController) EditDonation(ctx context.Context, in *donation_rest.EditDonationReq) (*donation_rest.EditResp, error) {

	return nil, nil
}

func (c *DonationController) DeleteDonation(ctx context.Context, in *donation_rest.DeleteDonationReq) (*donation_rest.DeleteResp, error) {

	return nil, nil
}
