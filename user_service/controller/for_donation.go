package controller

import (
	"context"
	"user_service/helper"
	"user_service/pb/user_donation"
)

func (c *UserDonation) GetUsername(ctx context.Context, in *user_donation.GetUsernameReq) (*user_donation.UsernameResp, error) {
	resp, err := c.UD.GetUsername(uint(in.UserId))
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	return &user_donation.UsernameResp{Username: resp}, nil
}
