package controller

import (
	"context"
	"registry_service/pb/pbDonationRegistry"
	"registry_service/pb/pbRegistryRest"
	"registry_service/repository"
)

type RegistryController struct {
	RR           repository.RegistryRepo
	DonationGRPC pbDonationRegistry.DonationRegistryClient
}

func (c *RegistryController) GetAllRegistries(ctx context.Context, in *pbRegistryRest.AllReq) (*pbRegistryRest.RegistriesResp, error) {

	return nil, nil
}

func (c *RegistryController) GetRegistryID(ctx context.Context, in *pbRegistryRest.GetRegistryReq) (*pbRegistryRest.DetailRegistryResp, error) {

	return nil, nil
}

func (c *RegistryController) Donate(ctx context.Context, in *pbRegistryRest.DonationReq) (*pbRegistryRest.DonateResp, error) {

	return nil, nil
}

func (c *RegistryController) DeleteRegistry(ctx context.Context, in *pbRegistryRest.DeleteRegistryReq) (*pbRegistryRest.DeleteResp, error) {

	return nil, nil
}
