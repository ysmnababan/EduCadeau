package handler

import (
	"api_gateway/helper"
	"api_gateway/pb"
	"api_gateway/pb/donation_rest"
	"api_gateway/pb/pbRegistryRest"
	"crypto/tls"
	"crypto/x509"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// InitUserHandler godoc
// @Summary Initialize User Handler
// @Description Initializes the gRPC client for the User service
// @Tags Initialization
// @Produce json
// @Success 200 {object} pb.UserToRestClient
// @Failure 500 {object} map[string]interface{}
// @Router /init/user [get]
func InitUserHandler() pb.UserToRestClient {
	// create connection to 'user service'
	addr := helper.USER_SERVICE_HOST + ":443"
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Fatalf("%s", err)
	}
	cred := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})

	// Initialize client connections outside handler in your implementation
	connection, err := grpc.Dial(addr, grpc.WithAuthority(helper.USER_SERVICE_HOST), grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Println(err)
	}
	userServiceClient := pb.NewUserToRestClient(connection)
	return userServiceClient
}

// InitDonationHandler godoc
// @Summary Initialize Donation Handler
// @Description Initializes the gRPC client for the Donation service
// @Tags Initialization
// @Produce json
// @Success 200 {object} donation_rest.DonationRestClient
// @Failure 500 {object} map[string]interface{}
// @Router /init/donation [get]
func InitDonationHandler() donation_rest.DonationRestClient {
	// create connection to 'donation service'
	addr := helper.DONATION_SERVICE_HOST + ":443"
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Fatalf("%s", err)
	}
	cred := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})

	// Initialize client connections outside handler in your implementation
	connection, err := grpc.Dial(addr, grpc.WithAuthority(helper.DONATION_SERVICE_HOST), grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Println(err)
	}
	donationServiceClient := donation_rest.NewDonationRestClient(connection)
	return donationServiceClient
}

// InitRegistryHandler godoc
// @Summary Initialize Registry Handler
// @Description Initializes the gRPC client for the Registry service
// @Tags Initialization
// @Produce json
// @Success 200 {object} pbRegistryRest.RegistryRestClient
// @Failure 500 {object} map[string]interface{}
// @Router /init/registry [get]
func InitRegistryHandler() pbRegistryRest.RegistryRestClient {
	// create connection to 'registry service'
	addr := helper.REGISTRY_SERVICE_HOST + ":443"
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Fatalf("%s", err)
	}
	cred := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})

	// Initialize client connections outside handler in your implementation
	connection, err := grpc.Dial(addr, grpc.WithAuthority(helper.REGISTRY_SERVICE_HOST), grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Println(err)
	}
	registryServiceClient := pbRegistryRest.NewRegistryRestClient(connection)
	return registryServiceClient
}
