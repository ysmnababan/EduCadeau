package setup

import (
	"fmt"
	"log"
	"net"
	"user_service/controller"
	"user_service/helper"
	"user_service/pb"
	"user_service/pb/user_donation"
	"user_service/pb/user_registry"

	"google.golang.org/grpc"
)

func SetupGPRCServer(UserRestServer *controller.UserController, UserDonationServer *controller.UserDonation, UserRegistryServer *controller.UserRegistry) {
	// create new grpc server
	grpcServer := grpc.NewServer()

	// register the 'user to rest' service server
	pb.RegisterUserToRestServer(grpcServer, UserRestServer)

	// register the 'user to donation' service server
	user_donation.RegisterUserDonationServer(grpcServer, UserDonationServer)

	// register the 'user to registry' service server
	user_registry.RegisterUserRegistryServer(grpcServer, UserRegistryServer)

	// start grpc server

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", helper.USER_SERVICE_PORT))
	if err != nil {
		log.Println(err)
	}

	log.Println("USER MICROSERVICE")
	log.Println("LISTENING TO PORT ", helper.USER_SERVICE_PORT)
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}
