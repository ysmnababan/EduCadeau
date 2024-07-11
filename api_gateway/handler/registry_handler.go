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

type RegistryHandler struct {
	RegistryGRPC pbRegistryRest.RegistryRestClient
}

func (h *RegistryHandler) GetAllRegistries(e echo.Context) error {
	cred := helper.GetCredential(e)
	if cred.Role == "recipient" {
		return helper.ParseError(helper.ErrDonorUser, e)
	}

	//get filter
	filter := e.QueryParam("filter")

	var donor_id uint64
	if cred.Role == "admin" {
		donor_id = 0
	} else {
		donor_id = uint64(cred.UserID)
	}
	res, err := h.RegistryGRPC.GetAllRegistries(context.TODO(), &pbRegistryRest.AllReq{Filter: filter, DonorId: donor_id})
	if err != nil {
		helper.Logging(e).Error("ERROR FROM REGISTRY GRPC: ", err)
		return helper.ParseErrorGRPC(err, e)
	}

	return e.JSON(http.StatusOK, res)
}

func (h *RegistryHandler) DetailOfRegistry(e echo.Context) error {
	cred := helper.GetCredential(e)
	if cred.Role != "donor" {
		return helper.ParseError(helper.ErrDonorUser, e)
	}

	//get id param
	registry_id := e.Param("id")
	if registry_id == "" {
		return helper.ParseError(helper.ErrInvalidId, e)
	}

	res, err := h.RegistryGRPC.GetRegistryID(context.TODO(), &pbRegistryRest.GetRegistryReq{RegistryId: registry_id, DonorId: uint64(cred.UserID)})

	if err != nil {
		helper.Logging(e).Error("ERROR FROM REGISTRY GRPC: ", err)
		return helper.ParseErrorGRPC(err, e)
	}

	return e.JSON(http.StatusOK, res)
}

func (h *RegistryHandler) Donate(e echo.Context) error {
	cred := helper.GetCredential(e)
	if cred.Role != "donor" {
		return helper.ParseError(helper.ErrDonorUser, e)
	}

	// get data bind
	var req models.CreateRegistryReq
	err := e.Bind(&req)
	if err != nil {
		log.Println("err bind: ", err, req)
		return helper.ParseError(helper.ErrBindJSON, e)
	}

	if req.DonationID == "" {
		return helper.ParseError(helper.ErrInvalidId, e)
	}

	// full
	if req.Filter == "" || (req.Filter != "full" && req.Filter != "partial") || req.Amount < 0 {
		return helper.ParseError(helper.ErrParam, e)
	}

	var amount float64
	if req.Filter == "partial" && req.Amount == 0 {
		return helper.ParseError(helper.ErrParam, e)
	}

	amount = 0.0
	if req.Filter == "partial" {
		amount = req.Amount
	}

	log.Println("req: ", req)
	res, err := h.RegistryGRPC.Donate(
		context.TODO(),
		&pbRegistryRest.DonationReq{
			DonationId: req.DonationID,
			Filter:     req.Filter,
			Amount:     amount,
			DonorId:    uint64(cred.UserID),
		},
	)
	if err != nil {
		helper.Logging(e).Error("ERROR FROM REGISTRY GRPC: ", err)
		return helper.ParseErrorGRPC(err, e)
	}

	return e.JSON(http.StatusCreated, res)
}

func (h *RegistryHandler) DeleteRegistry(e echo.Context) error {
	cred := helper.GetCredential(e)
	if cred.Role != "donor" {
		return helper.ParseError(helper.ErrDonorUser, e)
	}

	//get id param
	registry_id := e.Param("id")
	if registry_id == "" {
		return helper.ParseError(helper.ErrInvalidId, e)
	}

	res, err := h.RegistryGRPC.DeleteRegistry(
		context.TODO(),
		&pbRegistryRest.DeleteRegistryReq{
			DonorId: uint64(cred.UserID),
			RegistryId: registry_id,
		},
	)
	if err != nil {
		helper.Logging(e).Error("ERROR FROM REGISTRY GRPC: ", err)
		return helper.ParseErrorGRPC(err, e)
	}

	return e.JSON(http.StatusOK, res)
}

func (h *RegistryHandler) GetAllPayments(e echo.Context) error {

	return e.JSON(http.StatusOK, "")
}

func (h *RegistryHandler) GetPaymentID(e echo.Context) error {

	return e.JSON(http.StatusOK, "")
}

func (h *RegistryHandler) PayDonation(e echo.Context) error {

	return e.JSON(http.StatusOK, "")
}
