package controller

import (
	"context"
	"donation_service/helper"
	"donation_service/pb/pbDonationRegistry"
	"donation_service/pb/user_donation"
	"errors"
)

func (c *DonationController) GetDonationData(ctx context.Context, in *pbDonationRegistry.DonationReg) (*pbDonationRegistry.DonationResp, error) {
	res, err := c.DC.GetDonationDetail(in.DonationId)
	if err != nil {
		if errors.Is(err, helper.ErrNoData) {
			return &pbDonationRegistry.DonationResp{IsDonationExist: false}, err
		}
		return nil, err
	}
	user, err := c.UserGRPC.GetRecipientData(ctx, &user_donation.RecipientReq{UserId: uint64(res.RecipientID)})
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	return &pbDonationRegistry.DonationResp{
		IsDonationExist: true,
		DonationName:    res.DonationName,
		Description:     res.Description,
		RecipientId:     uint64(res.RecipientID),
		RecipientName:   user.Username,
		AmountToPay:     res.MiscellaneousCost + res.TargetAmount - res.AmountCollected,
	}, nil
}
