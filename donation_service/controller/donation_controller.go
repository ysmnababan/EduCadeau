package controller

import (
	"context"
	"donation_service/helper"
	"donation_service/models"
	"donation_service/pb/donation_rest"
	"donation_service/pb/user_donation"
	"donation_service/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DonationController struct {
	DC       repository.DonationRepo
	UserGRPC user_donation.UserDonationClient
}

func (c *DonationController) GetAllDonations(ctx context.Context, in *donation_rest.DonationReq) (*donation_rest.DonationList, error) {
	res, err := c.DC.GetAllDonations(in.Filter)
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}
	out := []*donation_rest.DonationDB{}

	for _, val := range res {
		d := &donation_rest.DonationDB{
			DonationId:        val.ID.String(),
			RecipientId:       uint64(val.RecipientID),
			DonationName:      val.DonationName,
			CreatedAt:         val.CreatedAt,
			Status:            val.Status,
			TargetAmount:      val.TargetAmount,
			AmountCollected:   val.AmountCollected,
			MiscellaneousCost: val.MiscellaneousCost,
		}
		out = append(out, d)
	}

	return &donation_rest.DonationList{List: out}, nil
}

func (c *DonationController) GetDonationDetail(ctx context.Context, in *donation_rest.DonationDetailReq) (*donation_rest.DonationDetailResp, error) {
	res, err := c.DC.GetDonationDetail(in.DonationId)
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}
	out := donation_rest.DonationDetailResp{
		DonationId:        res.ID.String(),
		RecipientId:       uint64(res.RecipientID),
		DonationName:      res.DonationName,
		CreatedAt:         res.CreatedAt,
		Status:            res.Status,
		TargetAmount:      res.TargetAmount,
		AmountCollected:   res.AmountCollected,
		MiscellaneousCost: res.MiscellaneousCost,
		Description:       res.Description,
		DonationType:      res.DonationType,
		Tag:               res.Tag,
		SenderAddress:     res.SenderAddress,
		RelatedLink:       res.RelatedLink,
		Notes:             res.Notes,
	}
	user, err := c.UserGRPC.GetRecipientData(ctx, &user_donation.RecipientReq{UserId: uint64(res.RecipientID)})
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}
	out.RecipientName = user.Username
	return &out, nil
}

func (c *DonationController) CreateDonation(ctx context.Context, in *donation_rest.CreateDonationReq) (*donation_rest.CreateResp, error) {
	user, err := c.UserGRPC.GetRecipientData(ctx, &user_donation.RecipientReq{UserId: uint64(in.RecipientId)})
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	miscCost := in.MiscellaneousCost
	if in.DonationType == "product" {
		miscCost, err = helper.GetPriceFromAddress(in.SenderAddress, user.Address)
		if err != nil {
			helper.Logging(nil).Error("ERROR GENERATING DELIVERYCOST: ", err)
			miscCost = 20000 // default delivery cost
		}
	}

	res, err := c.DC.CreateDonation(
		&models.CreateDonationReq{
			RecipientID:       uint(in.RecipientId),
			DonationName:      in.DonationName,
			TargetAmount:      in.TargetAmount,
			MiscellaneousCost: miscCost,
			Description:       in.Description,
			DonationType:      in.DonationType,
			Tag:               in.Tag,
			SenderAddress:     in.SenderAddress,
			RelatedLink:       in.RelatedLink,
			Notes:             in.Notes,
		},
	)
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	// send notification for user
	helper.CreateDonationNotif(&res)

	out := donation_rest.CreateResp{
		DonationId:        res.ID.String(),
		RecipientId:       uint64(res.RecipientID),
		DonationName:      res.DonationName,
		CreatedAt:         res.CreatedAt,
		Status:            res.Status,
		TargetAmount:      res.TargetAmount,
		AmountCollected:   res.AmountCollected,
		MiscellaneousCost: res.MiscellaneousCost,
		Description:       res.Description,
		DonationType:      res.DonationType,
		Tag:               res.Tag,
		SenderAddress:     res.SenderAddress,
		RelatedLink:       res.RelatedLink,
		Notes:             res.Notes,
		RecipientName:     user.Username,
	}

	return &out, nil
}

func (c *DonationController) EditDonation(ctx context.Context, in *donation_rest.EditDonationReq) (*donation_rest.EditResp, error) {
	user, err := c.UserGRPC.GetRecipientData(ctx, &user_donation.RecipientReq{UserId: uint64(in.RecipientId)})
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	d_id, err := primitive.ObjectIDFromHex(in.DonationId)
	if err != nil {
		return nil, helper.ParseErrorGRPC(helper.ErrInvalidId)
	}

	res, err := c.DC.EditDonation(
		&models.EditDonationReq{
			DonationID:    d_id,
			RecipientID:   uint(in.RecipientId),
			DonationName:  in.DonationName,
			TargetAmount:  in.TargetAmount,
			Description:   in.Description,
			DonationType:  in.DonationType,
			Tag:           in.Tag,
			SenderAddress: in.SenderAddress,
			RelatedLink:   in.RelatedLink,
			Notes:         in.Notes,
		},
	)
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	// send notification to user
	helper.EditDonationNotif(&res)

	out := donation_rest.EditResp{
		DonationId:        res.ID.String(),
		RecipientId:       uint64(res.RecipientID),
		DonationName:      res.DonationName,
		CreatedAt:         res.CreatedAt,
		Status:            res.Status,
		TargetAmount:      res.TargetAmount,
		AmountCollected:   res.AmountCollected,
		MiscellaneousCost: res.MiscellaneousCost,
		Description:       res.Description,
		DonationType:      res.DonationType,
		Tag:               res.Tag,
		SenderAddress:     res.SenderAddress,
		RelatedLink:       res.RelatedLink,
		Notes:             res.Notes,
		RecipientName:     user.Username,
	}

	return &out, nil
}

func (c *DonationController) DeleteDonation(ctx context.Context, in *donation_rest.DeleteDonationReq) (*donation_rest.DeleteResp, error) {

	res, err := c.DC.DeleteDonation(in.DonationId, uint(in.RecipientId))
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	return &donation_rest.DeleteResp{Message: res}, nil
}
