package setup

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"registry_service/helper"
	"registry_service/pb/pbDonationRegistry"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func SetupClientForDonationServer() pbDonationRegistry.DonationRegistryClient {
	// create connection to 'transaction service'
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
		log.Fatalf("did not connect: %v", err)
	}

	donationServiceClient := pbDonationRegistry.NewDonationRegistryClient(connection)

	return donationServiceClient
}
