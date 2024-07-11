package controller

import (
	"context"
	"log"
	"registry_service/helper"
	"registry_service/models"
	"registry_service/pb/pbDonationRegistry"
	"registry_service/pb/pbRegistryRest"
	"registry_service/pb/user_registry"
	"time"
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

	var out []*pbRegistryRest.PaymentResp
	for _, val := range resp {
		v := &pbRegistryRest.PaymentResp{
			PaymentAmount: val.PaymentAmount,
			PaymentDate:   val.PaymentDate,
			PaymentId:     val.ID.Hex(),
			PaymentMethod: val.PaymentMethod,
			RegistryId:    val.RegistryID.Hex(),
			InvoiceLink:   val.InvoiceLink,
		}
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
		InvoiceLink:   res.InvoiceLink,
	}, nil
}

func (c *RegistryController) Pay(ctx context.Context, in *pbRegistryRest.PayReq) (*pbRegistryRest.PaymentResp, error) {
	registryData, err := c.RR.GetRegistryID(in.RegistryId, in.DonorId)
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}
	log.Println("REGISTRY DATA: ", registryData)

	donationData, err := c.DonationGRPC.GetDonationData(
		context.TODO(),
		&pbDonationRegistry.DonationReg{
			DonationId: registryData.DonationID.Hex(),
		})
	if err != nil {
		helper.Logging(nil).Error("ERROR FROM DONATION GRPC: ", err)
		return nil, helper.ParseErrorGRPC(err)
	}
	log.Println("DONATION DATA: ", donationData)

	if donationData.Status == "settlement" || donationData.AmountToPay < registryData.Amount {
		// REGISTRY IS NO LONGER VALID, THERE IS CHANGE IN DONATION DATA
		// delete the registry
		c.RR.DeleteRegistry(registryData.ID.Hex(), in.DonorId)
		helper.Logging(nil).Error("REGISTRY IS NO LONGER VALID")
		return nil, helper.ParseErrorGRPC(helper.ErrInvalidRegistry)
	}

	user, err := c.UserGRPC.GetBalance(ctx, &user_registry.BalanceReq{UserId: in.DonorId})
	if err != nil {
		helper.Logging(nil).Error("ERROR FROM USER GRPC: ", err)
		return nil, helper.ParseErrorGRPC(err)
	}

	log.Println("DEPOSIT, AMOUNT : ", user.Deposit, registryData.Amount)

	var invoice_link string
	if in.PaymentMethod == "by deposit" {
		// using deposit
		if user.Deposit < registryData.Amount {
			return nil, helper.ParseErrorGRPC(helper.ErrUnsufficientBalance)
		}
		newDeposit := user.Deposit - registryData.Amount

		// grpc update user balance
		log.Println("update user balance: ", newDeposit, in.DonorId)
		_, err := c.UserGRPC.UpdateBalance(ctx, &user_registry.BalanceUpdate{UserId: in.DonorId, NewBalance: newDeposit})
		if err != nil {
			helper.Logging(nil).Error("ERROR FROM DONATION GRPC: ", err)
			return nil, helper.ParseErrorGRPC(err)
		}
	} else if in.PaymentMethod == "payment gateway" {
		res, err := helper.PaymentGateway(registryData.Amount, donationData)
		if err != nil {
			helper.Logging(nil).Error("ERROR GENERATING INVOICE: ", err)
			return nil, helper.ParseErrorGRPC(helper.ErrQuery)
		}
		invoice_link = res
	}

	//grpc update donation
	resp, err := c.DonationGRPC.AddAmountCollected(
		ctx,
		&pbDonationRegistry.AddReq{
			Amount:     registryData.Amount,
			DonationId: registryData.DonationID.Hex(),
		},
	)
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}
	log.Println("EDIT DONATION AFTER PAY: ", resp.AmountLeft, resp.Status)

	//create new payment record
	req := &models.Payment{
		RegistryID:    registryData.ID,
		PaymentDate:   time.Now().Format("2006-01-02 15:04:05"),
		PaymentMethod: in.PaymentMethod,
		PaymentAmount: registryData.Amount,
		InvoiceLink:   invoice_link,
	}
	log.Println("payment udpate data ", req)
	err = c.RR.Pay(req,
		in.DonorId)
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}
	return &pbRegistryRest.PaymentResp{
		PaymentId:     req.ID.Hex(),
		RegistryId:    req.RegistryID.Hex(),
		PaymentDate:   req.PaymentDate,
		PaymentMethod: req.PaymentMethod,
		PaymentAmount: req.PaymentAmount,
		InvoiceLink:   invoice_link,
	}, nil
}
