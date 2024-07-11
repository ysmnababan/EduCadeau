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

// GetAllDonations godoc
// @Summary Get all donations
// @Description Retrieve a list of all donations based on filters
// @Tags Donations
// @Produce json
// @Param filter query string false "Filter by donation status"
// @Success 200 {object} donation_rest.DonationResp
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /donations [get]
func (h *DonationHandler) GetAllDonations(e echo.Context) error {
	cred := helper.GetCredential(e)

	filter := e.QueryParam("filter")
	if filter == "settled" && cred.Role != "admin" {
		return helper.ParseError(helper.ErrMustAdmin, e)
	}
	if (filter == "on progress" || filter == "unsponsored") && cred.Role == "recipient" {
		return helper.ParseError(helper.ErrDonorUser, e)
	}
	if filter == "requested" && cred.Role != "recipient" {
		return helper.ParseError(helper.ErrRecipientUser, e)
	}

	resp, err := h.DonationGRPC.GetAllDonations(e.Request().Context(), &donation_rest.DonationReq{Filter: filter})
	if err != nil {
		helper.Logging(e).Error("FROM GRPC: ", err)
		return helper.ParseErrorGRPC(err, e)
	}
	return e.JSON(http.StatusOK, resp)
}

// GetDonationDetail godoc
// @Summary Get donation details
// @Description Retrieve the details of a specific donation
// @Tags Donations
// @Produce json
// @Param id path string true "Donation ID"
// @Success 200 {object} donation_rest.DonationDetailResp
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /donation/{id} [get]
func (h *DonationHandler) GetDonationDetail(e echo.Context) error {

	donation_id := e.Param("id")

	if donation_id == "" {
		return helper.ParseError(helper.ErrInvalidId, e)
	}

	resp, err := h.DonationGRPC.GetDonationDetail(e.Request().Context(), &donation_rest.DonationDetailReq{DonationId: donation_id})
	if err != nil {
		helper.Logging(e).Error("FROM GRPC: ", err)
		return helper.ParseErrorGRPC(err, e)
	}

	return e.JSON(http.StatusOK, resp)
}

// CreateDonation godoc
// @Summary Create a new donation
// @Description Create a new donation request
// @Tags Donations
// @Accept json
// @Produce json
// @Param donation body models.CreateDonationReq true "Donation request payload"
// @Success 201 {object} donation_rest.CreateDonationResp
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /donation [post]
func (h *DonationHandler) CreateDonation(e echo.Context) error {
	cred := helper.GetCredential(e)
	if cred.Role != "recipient" {
		return helper.ParseError(helper.ErrRecipientUser, e)
	}

	var in models.CreateDonationReq
	err := e.Bind(&in)
	if err != nil {
		helper.Logging(e).Error("ERR BIND: ", err)
		return helper.ParseError(helper.ErrBindJSON, e)
	}

	//validate
	if in.DonationName == "" || in.TargetAmount <= 0 || in.Description == "" || in.DonationType != "service" && in.DonationType != "product" || in.MiscellaneousCost < 0 {
		return helper.ParseError(helper.ErrParam, e)
	}

	// if donation type is "product",  address should not be empty,
	// miscellaneous cost cannot be inputted manually, calculated using third party api
	if in.DonationType == "product" && in.SenderAddress == "" {
		helper.Logging(e).Error("product type must have sender address")
		return helper.ParseError(helper.ErrParam, e)
	}

	// if donation type is "service", address can be empty and miscellaneous cost can input manually

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

// EditDonation godoc
// @Summary Edit a donation
// @Description Edit an existing donation
// @Tags Donations
// @Accept json
// @Produce json
// @Param id path string true "Donation ID"
// @Param donation body models.EditDonationReq true "Donation edit payload"
// @Success 200 {object} donation_rest.EditDonationResp
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /donation/{id} [put]
func (h *DonationHandler) EditDonation(e echo.Context) error {
	cred := helper.GetCredential(e)
	if cred.Role != "recipient" {
		return helper.ParseError(helper.ErrRecipientUser, e)
	}

	var in models.EditDonationReq
	err := e.Bind(&in)
	if err != nil {
		helper.Logging(e).Error("ERR BIND: ", err)
		return helper.ParseError(helper.ErrBindJSON, e)
	}

	donation_id := e.Param("id")
	if donation_id == "" {
		return helper.ParseError(helper.ErrInvalidId, e)
	}

	//validate
	if in.TargetAmount <= 0 || in.MiscellaneousCost < 0 {
		return helper.ParseError(helper.ErrParam, e)
	}

	resp, err := h.DonationGRPC.EditDonation(
		context.TODO(),
		&donation_rest.EditDonationReq{
			DonationId:        donation_id,
			RecipientId:       uint64(cred.UserID),
			DonationName:      in.DonationName,
			TargetAmount:      in.TargetAmount,
			Description:       in.Description,
			Tag:               in.Tag,
			SenderAddress:     in.SenderAddress,
			RelatedLink:       in.RelatedLink,
			Notes:             in.Notes,
			MiscellaneousCost: in.MiscellaneousCost,
		},
	)
	if err != nil {
		helper.Logging(e).Error("FROM GRPC: ", err)
		return helper.ParseErrorGRPC(err, e)
	}
	log.Println("EDIT RESP: ", resp)
	return e.JSON(http.StatusOK, resp)
}

// DeleteDonation godoc
// @Summary Delete a donation
// @Description Delete an existing donation
// @Tags Donations
// @Produce json
// @Param id path string true "Donation ID"
// @Success 200 {object} donation_rest.DeleteDonationResp
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /donation/{id} [delete]
func (h *DonationHandler) DeleteDonation(e echo.Context) error {
	cred := helper.GetCredential(e)
	if cred.Role != "recipient" {
		return helper.ParseError(helper.ErrRecipientUser, e)
	}

	donation_id := e.Param("id")
	if donation_id == "" {
		return helper.ParseError(helper.ErrInvalidId, e)
	}

	resp, err := h.DonationGRPC.DeleteDonation(
		context.TODO(),
		&donation_rest.DeleteDonationReq{
			DonationId:  donation_id,
			RecipientId: uint64(cred.UserID),
		},
	)
	if err != nil {
		helper.Logging(e).Error("FROM GRPC: ", err)
		return helper.ParseErrorGRPC(err, e)
	}
	return e.JSON(http.StatusOK, resp)
}
