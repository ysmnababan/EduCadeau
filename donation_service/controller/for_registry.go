package controller

import (
	"context"
	"donation_service/helper"
	"donation_service/pb/pbDonationRegistry"
	"donation_service/pb/user_donation"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (c *DonationController) AddAmountCollected(ctx context.Context, in *pbDonationRegistry.AddReq) (*pbDonationRegistry.AddResp, error) {
	donation_id, err := primitive.ObjectIDFromHex(in.DonationId)
	if err != nil {
		helper.Logging(nil).Error("ERR REPO: ", err)
		return nil, helper.ErrInvalidId
	}

	res, err := c.DC.EditDonationAfterPay(donation_id, in.Amount)
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	return &pbDonationRegistry.AddResp{
		AmountLeft: res.AmountLeft,
		Status:     res.Status,
	}, nil
}
