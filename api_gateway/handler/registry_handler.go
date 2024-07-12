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

// GetAllRegistries godoc
// @Summary Get all registries
// @Description Get all registries for a user
// @Tags Registry
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentication token"
// @Param filter query string false "Filter by registry status"
// @Success 200 {array} pbRegistryRest.RegistriesResp
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /donated [get]
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

// DetailOfRegistry godoc
// @Summary Get registry detail
// @Description Get details of a specific registry
// @Tags Registry
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentication token"
// @Param id path string true "Registry ID"
// @Success 200 {object} pbRegistryRest.DetailRegistryResp
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /donated/{id} [get]
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

// Donate godoc
// @Summary Create a donation registry
// @Description Create a new donation registry
// @Tags Registry
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentication token"
// @Param donation body models.CreateRegistryReq true "Donation request data"
// @Success 201 {object} pbRegistryRest.DonateResp
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /donate [post]
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

// DeleteRegistry godoc
// @Summary Delete a registry
// @Description Delete a registry by ID
// @Tags Registry
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentication token"
// @Param id path string true "Registry ID"
// @Success 200 {object} pbRegistryRest.DeleteResp
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /donated/{id} [delete]
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
			DonorId:    uint64(cred.UserID),
			RegistryId: registry_id,
		},
	)
	if err != nil {
		helper.Logging(e).Error("ERROR FROM REGISTRY GRPC: ", err)
		return helper.ParseErrorGRPC(err, e)
	}

	return e.JSON(http.StatusOK, res)
}
