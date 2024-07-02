package setup

import (
	"fmt"
	"log"
	"net"
	"user_service/controller"
	"user_service/helper"
	"user_service/pb"

	"google.golang.org/grpc"
)

func SetupGPRCServer(UC *controller.UserController) {
	// create new grpc server
	grpcServer := grpc.NewServer()

	// register the 'payment' service server
	pb.RegisterUserToRestServer(grpcServer, UC)
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
