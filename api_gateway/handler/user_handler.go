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

// Register godoc
// @Summary Register as user
// @Description Register as user and return user data
// @Tags User
// @Accept  json
// @Produce  json
// @Param student body models.UserRegister true "Register new user"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users/register [post]
func (h *UserHandler) Register(c echo.Context) error {
	var GetU models.User
	err := c.Bind(&GetU)
	if err != nil {
		return helper.ParseError(helper.ErrBindJSON, c)
	}

	//validate user
	if GetU.Email == "" || GetU.Password == "" || GetU.Username == "" {
		return helper.ParseError(helper.ErrParam, c)
	}

	// validate role
	if GetU.Role != "admin" && GetU.Role != "recipient" && GetU.Role != "donor" {
		return helper.ParseError(helper.ErrParam, c)
	}

	respU, err := h.UserGRPC.Register(
		context.TODO(),
		&pb.RegisterReq{
			Username: GetU.Username,
			Password: GetU.Password,
			Email:    GetU.Email,
			Role:     GetU.Role,
		},
	)
	if err != nil {
		return helper.ParseErrorGRPC(err, c)
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "New User Created", "User": respU})
}

// GetAllUser godoc
// @Summary Get info about a user ONLY FOR ADMIN
// @Description must be authenticated user and return all user data
// @Tags User
// @Accept  json
// @Produce  json
// @Param   Authorization  header  string  true  "Authentication token"  default()
// @Success 200 {array} models.UserDetailResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users [get]
func (h *UserHandler) GetAllUser(c echo.Context) error {
	cred := helper.GetCredential(c)
	if cred.Role != "admin" {
		return helper.ParseError(helper.ErrMustAdmin, c)
	}

	resp, err := h.UserGRPC.GetAllUser(context.TODO(), &pb.Req{})
	if err != nil {
		return helper.ParseError(err, c)
	}
	log.Println("get all user: ", resp.List)
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Get All User", "User": resp.List})
}

// GetUserInfo godoc
// @Summary Get info about a user
// @Description must be authenticated user and return user detail data
// @Tags User
// @Accept  json
// @Produce  json
// @Param   Authorization  header  string  true  "Authentication token"  default()
// @Success 200 {object} models.UserDetailResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/user [get]
func (h *UserHandler) GetUserDetail(c echo.Context) error {
	cred := helper.GetCredential(c)
	resp, err := h.UserGRPC.GetUserDetail(
		context.TODO(),
		&pb.DetailReq{UserId: uint64(cred.UserID)},
	)
	if err != nil {
		return helper.ParseErrorGRPC(err, c)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Get User Info", "User": resp})
}

func (h *UserHandler) EditUser(c echo.Context) error {

	return nil
}

func (h *UserHandler) TopUp(c echo.Context) error {

	return nil
}
