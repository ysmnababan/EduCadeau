package handler

import (
	"api_gateway/helper"
	"api_gateway/models"
	"api_gateway/pb"
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"

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

// Login godoc
// @Summary Login as user
// @Description login as user and generate token
// @Tags User
// @Accept  json
// @Produce  json
// @Param student body models.UserRequest true "Login using email and password"
// @Success 200 {object} map[string]interface{} "message : string, token: string"
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users/login [post]
func (h *UserHandler) Login(c echo.Context) error {
	var GetU models.User
	err := c.Bind(&GetU)
	if err != nil {
		return helper.ParseError(helper.ErrBindJSON, c)
	}

	//validate user
	if GetU.Email == "" || GetU.Password == "" {
		return helper.ParseError(helper.ErrParam, c)
	}

	tokenString, err := h.UserGRPC.Login(
		context.TODO(),
		&pb.LoginReq{
			Email:    GetU.Email,
			Password: GetU.Password,
		})
	if err != nil {
		return helper.ParseErrorGRPC(err, c)
	}

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"message": "Login success",
			"token":   tokenString,
		})
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
