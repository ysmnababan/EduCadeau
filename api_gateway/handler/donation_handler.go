package handler

import (
	"api_gateway/helper"
	"api_gateway/pb/donation_rest"
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
		return helper.ParseErrorGRPC(err, e)
	}
	return e.JSON(http.StatusOK, resp)
}

func (h *DonationHandler) GetDonationDetail(e echo.Context) error {
	type req struct {
		DonationID string `json:"donation_id" bson:"donation_id"`
	}

	var idreq req
	err := e.Bind(&idreq)
	if err != nil {
		helper.Logging(e).Error("ERR BIND: ", err)
		return helper.ParseError(helper.ErrBindJSON, e)
	}
	log.Println(idreq)

	resp, err := h.DonationGRPC.GetDonationDetail(e.Request().Context(), &donation_rest.DonationDetailReq{DonationId: idreq.DonationID})
	if err != nil {
		return helper.ParseErrorGRPC(err, e)
	}

	return e.JSON(http.StatusOK, resp)
}

func (h *DonationHandler) CreateDonation(e echo.Context) error {
	resp := ""
	return e.JSON(http.StatusOK, resp)
}

func (h *DonationHandler) EditDonation(e echo.Context) error {
	resp := ""
	return e.JSON(http.StatusOK, resp)
}

func (h *DonationHandler) DeleteDonation(e echo.Context) error {
	resp := ""
	return e.JSON(http.StatusOK, resp)
}
