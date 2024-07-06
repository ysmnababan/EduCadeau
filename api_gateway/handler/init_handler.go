package handler

import (
	"api_gateway/helper"
	"api_gateway/pb"
	"api_gateway/pb/donation_rest"
	"crypto/tls"
	"crypto/x509"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

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
