package handler

import (
	"api_gateway/helper"
	"api_gateway/models"
	"api_gateway/pb/donation_rest"
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DonationHandler struct {
	DonationGRPC donation_rest.DonationRestClient
}

func (h *DonationHandler) GetAllDonations(e echo.Context) error {
	cred := helper.GetCredential(e)

	filter := e.QueryParam("filter")
	if filter == "settled" && cred.Role != "admin" {
		return e.JSON(http.StatusUnauthorized, helper.ErrMustAdmin)
	}
	if (filter == "on progress" || filter == "unsponsored") && cred.Role == "recipient" {
		return e.JSON(http.StatusUnauthorized, helper.ErrDonorUser)
	}
	if filter == "requested" && cred.Role != "recipient" {
		return e.JSON(http.StatusUnauthorized, helper.ErrRecipientUser)
	}

	resp, err := h.DonationGRPC.GetAllDonations(e.Request().Context(), &donation_rest.DonationReq{Filter: filter})
	if err != nil {
		helper.Logging(e).Error("FROM GRPC: ", err)
		return helper.ParseErrorGRPC(err, e)
	}
	return e.JSON(http.StatusOK, resp)
}

func (h *DonationHandler) GetDonationDetail(e echo.Context) error {

	var idreq models.ReqID
	err := e.Bind(&idreq)
	if err != nil {
		helper.Logging(e).Error("ERR BIND: ", err)
		return helper.ParseError(helper.ErrBindJSON, e)
	}
	log.Println(idreq)

	if idreq.DonationID == "" {
		return e.JSON(http.StatusBadGateway, helper.ErrInvalidId)

	}

	resp, err := h.DonationGRPC.GetDonationDetail(e.Request().Context(), &donation_rest.DonationDetailReq{DonationId: idreq.DonationID})
	if err != nil {
		helper.Logging(e).Error("FROM GRPC: ", err)
		return helper.ParseErrorGRPC(err, e)
	}

	return e.JSON(http.StatusOK, resp)
}

func (h *DonationHandler) CreateDonation(e echo.Context) error {
	cred := helper.GetCredential(e)
	if cred.Role != "recipient" {
		return e.JSON(http.StatusUnauthorized, helper.ErrRecipientUser)
	}
	var in models.CreateDonationReq
	err := e.Bind(&in)
	if err != nil {
		helper.Logging(e).Error("ERR BIND: ", err)
		return helper.ParseError(helper.ErrBindJSON, e)
	}

	//validate
	if in.DonationName == "" || in.TargetAmount <= 0 || in.Description == "" || in.DonationType != "service" && in.DonationType != "product" || in.MiscellaneousCost < 0 || int(in.RecipientID) < 0 {
		return e.JSON(http.StatusBadGateway, helper.ErrParam)
	}
	resp, err := h.DonationGRPC.CreateDonation(
		context.TODO(),
		&donation_rest.CreateDonationReq{
			RecipientId:       uint64(cred.UserID),
			DonationName:      in.DonationName,
			TargetAmount:      in.TargetAmount,
			MiscellaneousCost: in.MiscellaneousCost,
			Description:       in.Description,
			DonationType:      in.DonationType,
			Tag:               in.Tag,
			SenderAddress:     in.SenderAddress,
			RelatedLink:       in.RelatedLink,
			Notes:             in.Notes,
		},
	)
	if err != nil {
		helper.Logging(e).Error("FROM GRPC: ", err)
		return helper.ParseErrorGRPC(err, e)
	}
	log.Println("CREATE RESP: ", resp)

	return e.JSON(http.StatusCreated, resp)
}

func (h *DonationHandler) EditDonation(e echo.Context) error {
	cred := helper.GetCredential(e)
	if cred.Role != "recipient" {
		return e.JSON(http.StatusUnauthorized, helper.ErrRecipientUser)
	}
	var in models.EditDonationReq
	err := e.Bind(&in)
	if err != nil {
		helper.Logging(e).Error("ERR BIND: ", err)
		return helper.ParseError(helper.ErrBindJSON, e)
	}

	//validate
	if in.DonationID == "" || in.DonationName == "" || in.TargetAmount <= 0 || in.Description == "" || in.DonationType != "service" && in.DonationType != "product" || int(in.RecipientID) < 0 {
		return e.JSON(http.StatusBadGateway, helper.ErrParam)
	}
	resp, err := h.DonationGRPC.EditDonation(
		context.TODO(),
		&donation_rest.EditDonationReq{
			DonationId:    in.DonationID,
			RecipientId:   uint64(cred.UserID),
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
		helper.Logging(e).Error("FROM GRPC: ", err)
		return helper.ParseErrorGRPC(err, e)
	}
	log.Println("EDIT RESP: ", resp)
	return e.JSON(http.StatusOK, resp)
}

func (h *DonationHandler) DeleteDonation(e echo.Context) error {
	cred := helper.GetCredential(e)
	if cred.Role != "recipient" {
		return e.JSON(http.StatusUnauthorized, helper.ErrRecipientUser)
	}

	var idreq models.ReqID
	err := e.Bind(&idreq)
	if err != nil {
		helper.Logging(e).Error("ERR BIND: ", err)
		return helper.ParseError(helper.ErrBindJSON, e)
	}

	log.Println(idreq)

	if idreq.DonationID == "" {
		return e.JSON(http.StatusBadGateway, helper.ErrInvalidId)

	}

	resp, err := h.DonationGRPC.DeleteDonation(
		context.TODO(),
		&donation_rest.DeleteDonationReq{
			DonationId:  idreq.DonationID,
			RecipientId: uint64(cred.UserID),
		},
	)
	if err != nil {
		helper.Logging(e).Error("FROM GRPC: ", err)
		return helper.ParseErrorGRPC(err, e)
	}
	return e.JSON(http.StatusOK, resp)
}
