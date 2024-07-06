package setup

import (
	"crypto/tls"
	"crypto/x509"
	"donation_service/helper"
	"donation_service/pb/user_donation"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func SetupClientForUserServer() user_donation.UserDonationClient {
	// create connection to 'transaction service'
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
		log.Fatalf("did not connect: %v", err)
	}

	donationServiceClient := user_donation.NewUserDonationClient(connection)

	return donationServiceClient
}
