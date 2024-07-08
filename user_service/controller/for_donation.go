package controller

import (
	"context"
	"user_service/helper"
	"user_service/pb/user_donation"
)

func (c *UserDonation) GetRecipientData(ctx context.Context, in *user_donation.RecipientReq) (*user_donation.DetailResp, error) {
	resp, err := c.UD.GetInfo(uint(in.UserId))
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	return &user_donation.DetailResp{Username: resp.Username, Address: resp.Address}, nil
}
