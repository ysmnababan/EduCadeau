package handler

import (
	"api_gateway/helper"
	"api_gateway/pb"
	"crypto/tls"
	"crypto/x509"
	"log"

	"github.com/labstack/echo/v4"
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

type UserHandler struct {
	UserGRPC pb.UserToRestClient
}

func (h *UserHandler) Login(e echo.Context) error {

	return nil
}

func (h *UserHandler) Register(e echo.Context) error {

	return nil
}

func (h *UserHandler) GetAllUser(e echo.Context) error {

	return nil
}

func (h *UserHandler) GetUserDetail(e echo.Context) error {

	return nil
}

func (h *UserHandler) EditUser(e echo.Context) error {

	return nil
}

func (h *UserHandler) TopUp(e echo.Context) error {

	return nil
}
