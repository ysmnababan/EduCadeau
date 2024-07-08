package controller

import (
	"context"
	"log"
	"registry_service/helper"
	"registry_service/models"
	"registry_service/pb/pbRegistryRest"
)

func (c *RegistryController) GetAllPayments(ctx context.Context, in *pbRegistryRest.PaymentsReq) (*pbRegistryRest.PaymentList, error) {
	// var resp :=0
	var resp []*models.Payment
	var err error
	if in.DonorId == 0 {
		resp, err = c.RR.GetAllPayments()
	} else {
		resp, err = c.RR.GetAllMyPayments(in.DonorId)
	}

	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	out := []*pbRegistryRest.PaymentResp{}
	for _, val := range resp {
		var v *pbRegistryRest.PaymentResp
		v.PaymentAmount = val.PaymentAmount
		v.PaymentDate = val.PaymentDate
		v.PaymentId = val.ID.Hex()
		v.PaymentMethod = val.PaymentMethod
		v.RegistryId = val.RegistryID.Hex()

		out = append(out, v)
	}

	log.Printf("GET ALL PAYMENT: %v\n+++++++++++++++++\n\n\n", out)
	return &pbRegistryRest.PaymentList{List: out}, nil
}

func (c *RegistryController) GetPayment(ctx context.Context, in *pbRegistryRest.PaymentReq) (*pbRegistryRest.PaymentResp, error) {
	res, err := c.RR.GetPayment(in.DonorId, in.PaymentId)
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	log.Printf("GET PAYMENT: %v\n+++++++++++++++++\n\n\n", res)
	return &pbRegistryRest.PaymentResp{
		PaymentId:     res.ID.Hex(),
		RegistryId:    res.RegistryID.Hex(),
		PaymentDate:   res.PaymentDate,
		PaymentMethod: res.PaymentMethod,
		PaymentAmount: res.PaymentAmount,
	}, nil
}
