package handler

import (
	"api_gateway/pb/pbRegistryRest"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RegistryHandler struct {
	RegistryGRPC pbRegistryRest.RegistryRestClient
}

func (h *RegistryHandler) GetAllRegistries(e echo.Context) error {

	return e.JSON(http.StatusOK, "")
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
