package handler

import (
	"api_gateway/helper"
	"api_gateway/models"
	"api_gateway/pb/pbRegistryRest"
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *RegistryHandler) GetAllPayments(e echo.Context) error {
	cred := helper.GetCredential(e)
	if cred.Role == "recipient" {
		return helper.ParseError(helper.ErrDonorUser, e)
	}

	//get filter
	donor_id := 0
	if cred.Role == "donor" {
		donor_id = int(cred.UserID)
	}
	log.Println(donor_id)
	res, err := h.RegistryGRPC.GetAllPayments(context.TODO(), &pbRegistryRest.PaymentsReq{DonorId: uint64(donor_id)})
	if err != nil {
		helper.Logging(e).Error("ERROR FROM REGISTRY GRPC: ", err)
		return helper.ParseErrorGRPC(err, e)
	}

	return e.JSON(http.StatusOK, res)
}

func (h *RegistryHandler) GetPaymentID(e echo.Context) error {
	cred := helper.GetCredential(e)
	if cred.Role != "donor" {
		return helper.ParseError(helper.ErrDonorUser, e)
	}
	payment_id := e.Param("id")
	if payment_id == "" {
		return helper.ParseError(helper.ErrInvalidId, e)
	}

	res, err := h.RegistryGRPC.GetPayment(context.TODO(), &pbRegistryRest.PaymentReq{DonorId: uint64(cred.UserID), PaymentId: payment_id})
	if err != nil {
		helper.Logging(e).Error("ERROR FROM REGISTRY GRPC: ", err)
		return helper.ParseErrorGRPC(err, e)
	}

	return e.JSON(http.StatusOK, res)
}

func (h *RegistryHandler) PayDonation(e echo.Context) error {
	cred := helper.GetCredential(e)
	if cred.Role != "donor" {
		return helper.ParseError(helper.ErrDonorUser, e)
	}

	req := models.PayReq{}
	err := e.Bind(&req)
	if err != nil {
		log.Println("err bind: ", err, req)
		return helper.ParseError(helper.ErrBindJSON, e)
	}

	if req.PaymentMethod == "" || (req.PaymentMethod != "by deposit" && req.PaymentMethod != "payment gateway") {
		return helper.ParseError(helper.ErrParam, e)
	}

	registry_id := e.Param("id")
	if registry_id == "" {
		return helper.ParseError(helper.ErrInvalidId, e)
	}

	resp, err := h.RegistryGRPC.Pay(
		context.TODO(),
		&pbRegistryRest.PayReq{
			DonorId:       uint64(cred.UserID),
			RegistryId:    registry_id,
			PaymentMethod: req.PaymentMethod,
		},
	)

	if err != nil {
		helper.Logging(e).Error("ERROR FROM REGISTRY GRPC: ", err)
		return helper.ParseErrorGRPC(err, e)
	}

	return e.JSON(http.StatusCreated, resp)
}
