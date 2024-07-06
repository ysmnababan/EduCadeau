package handler

import (
	"api_gateway/pb/donation_rest"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DonationHandler struct {
	DonationGRPC donation_rest.DonationRestClient
}

func (h *DonationHandler) GetAllDonations(e echo.Context) error {

	resp := ""
	return e.JSON(http.StatusOK, resp)
}

func (h *DonationHandler) GetDonationDetail(e echo.Context) error {
	resp := ""
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
