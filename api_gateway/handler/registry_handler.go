package handler

import (
	"api_gateway/helper"
	"api_gateway/pb/pbRegistryRest"
	"context"
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

	return e.JSON(http.StatusOK, "")
}

func (h *RegistryHandler) Donate(e echo.Context) error {

	return e.JSON(http.StatusOK, "")
}

func (h *RegistryHandler) DeleteRegistry(e echo.Context) error {

	return e.JSON(http.StatusOK, "")
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
