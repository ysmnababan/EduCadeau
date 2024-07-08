package setup

import (
	"fmt"
	"log"
	"net"
	"registry_service/controller"
	"registry_service/helper"
	"registry_service/pb/pbRegistryRest"

	"google.golang.org/grpc"
)

func SetupGPRCServer(RC *controller.RegistryController) {
	// create new grpc server
	grpcServer := grpc.NewServer()

	// register the 'Donation' service server
	pbRegistryRest.RegisterRegistryRestServer(grpcServer, RC)
	// start grpc server

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", helper.REGISTRY_SERVICE_PORT))
	if err != nil {
		log.Println(err)
	}

	log.Println("REGISTRY MICROSERVICE")
	log.Println("LISTENING TO PORT ", helper.REGISTRY_SERVICE_PORT)
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}
