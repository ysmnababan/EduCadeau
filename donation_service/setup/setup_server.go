package setup

import (
	"donation_service/controller"
	"donation_service/helper"
	"donation_service/pb/donation_rest"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func SetupGPRCServer(DC *controller.DonationController) {
	// create new grpc server
	grpcServer := grpc.NewServer()

	// register the 'Donation' service server
	donation_rest.RegisterDonationRestServer(grpcServer, DC)
	// start grpc server

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", helper.DONATION_SERVICE_PORT))
	if err != nil {
		log.Println(err)
	}

	log.Println("DONATION MICROSERVICE")
	log.Println("LISTENING TO PORT ", helper.DONATION_SERVICE_PORT)
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}
