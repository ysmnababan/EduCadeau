package controller

import (
	"context"
	"log"
	"registry_service/helper"
	"registry_service/models"
	"registry_service/pb/pbDonationRegistry"
	"registry_service/pb/pbRegistryRest"
	"registry_service/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegistryController struct {
	RR           repository.RegistryRepo
	DonationGRPC pbDonationRegistry.DonationRegistryClient
}

func (c *RegistryController) GetAllRegistries(ctx context.Context, in *pbRegistryRest.AllReq) (*pbRegistryRest.RegistriesResp, error) {
	var out []*pbRegistryRest.RegistryResp
	res, err := c.RR.GetAllRegistries(in.Filter)
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}
	for _, val := range res {
		o := &pbRegistryRest.RegistryResp{
			RegistryId: val.ID.String(),
			DonationId: val.DonationID.String(),
			DonorId:    val.DonorID,
			Amount:     val.Amount,
			Status:     val.Status,
		}
		out = append(out, o)
	}

	log.Printf("GET ALL REGISTRY SUCCESS DATA: %v\n\n", out)
	return &pbRegistryRest.RegistriesResp{List: out}, nil
}

func (c *RegistryController) GetRegistryID(ctx context.Context, in *pbRegistryRest.GetRegistryReq) (*pbRegistryRest.DetailRegistryResp, error) {
	res, err := c.RR.GetRegistryID(in.RegistryId, in.DonorId)
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	d, err := c.DonationGRPC.GetDonationData(
		ctx,
		&pbDonationRegistry.DonationReg{
			DonationId: res.DonationID.String(),
		},
	)

	if err != nil {
		helper.Logging(nil).Error("ERROR FROM DONATION GRPC: ", err)
	}

	helper.Logging(nil).Error("REGISTRY DATA: ", res)
	helper.Logging(nil).Error("DONATION DATA: ", d)

	log.Printf("GET REGISTRY SUCCESS DATA: \n\n")

	return &pbRegistryRest.DetailRegistryResp{
		RegistryId:   in.RegistryId,
		DonationId:   res.DonationID.String(),
		DonorId:      in.DonorId,
		Amount:       res.Amount,
		Status:       res.Status,
		DonationName: d.DonationName,
		Description:  d.Description,
		AmountToPay:  d.AmountToPay,
		RecipientId:  d.RecipientId,
	}, nil
}

func (c *RegistryController) Donate(ctx context.Context, in *pbRegistryRest.DonationReq) (*pbRegistryRest.DonateResp, error) {
	d_id, err := primitive.ObjectIDFromHex(in.DonationId)
	if err != nil {
		helper.Logging(nil).Error("ERR REPO: ", err)
		return nil, helper.ParseErrorGRPC(helper.ErrInvalidId)
	}

	d, err := c.DonationGRPC.GetDonationData(
		ctx,
		&pbDonationRegistry.DonationReg{
			DonationId: in.DonationId,
		},
	)

	if err != nil {
		helper.Logging(nil).Error("ERROR FROM DONATION GRPC: ", err)
		return nil, helper.ParseErrorGRPC(helper.ErrQuery)
	}

	helper.Logging(nil).Error("DONATION DATA: ", d)

	if in.Filter == "full" {
		in.Amount = d.AmountToPay
	} else if in.Filter == "partial" && in.Amount >= d.AmountToPay {
		return nil, helper.ParseErrorGRPC(helper.ErrParam)
	}
	req := &models.Registry{
		DonationID: d_id,
		DonorID:    in.DonorId,
		Amount:     in.Amount,
	}
	err = c.RR.Donate(req)

	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	log.Printf("CREATE DONATE SUCCESS DATA: %v\n\n", req)

	return &pbRegistryRest.DonateResp{
		RegistryId: req.ID.String(),
		DonationId: req.DonationID.String(),
		DonorId:    req.DonorID,
		Amount:     req.Amount,
		Status:     req.Status,
	}, nil
}

func (c *RegistryController) DeleteRegistry(ctx context.Context, in *pbRegistryRest.DeleteRegistryReq) (*pbRegistryRest.DeleteResp, error) {
	res, err := c.RR.DeleteRegistry(in.RegistryId, in.DonorId)
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	log.Printf("DELETE REGISTRY SUCCESS DATA: %v\n\n", res)

	return &pbRegistryRest.DeleteResp{Message: res}, nil
}
