package handler

import (
	"api_gateway/helper"
	"api_gateway/models"
	"api_gateway/pb"
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
			"token":   tokenString.Token,
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

// UpdateUser godoc
// @Summary Update user information
// @Description must be authenticated user and update detail info of a user
// @Tags User
// @Accept  json
// @Produce  json
// @Param   Authorization  header  string  true  "Authentication token"  default()
// @Param student body models.UserUpdateRequest true "Data to be updated"
// @Success 200 {object} models.UserDetailResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/user [put]
func (h *UserHandler) EditUser(c echo.Context) error {
	cred := helper.GetCredential(c)
	var GetU models.UserUpdateRequest
	err := c.Bind(&GetU)
	if err != nil {
		log.Println("ERROR BINDING: ", err)
		return helper.ParseError(helper.ErrBindJSON, c)
	}

	//validate user
	if GetU.Username == "" || GetU.Age <= 0 {
		return helper.ParseError(helper.ErrParam, c)
	}

	respU, err := h.UserGRPC.EditDataUser(
		context.TODO(),
		&pb.EditReq{
			Username:          GetU.Username,
			Fname:             GetU.Fname,
			Lname:             GetU.Lname,
			Address:           GetU.Address,
			Age:               int64(GetU.Age),
			PhoneNumber:       GetU.PhoneNumber,
			ProfilePictureUrl: GetU.ProfilePictureUrl,
			UserId:            uint64(cred.UserID),
		},
	)
	if err != nil {
		return helper.ParseErrorGRPC(err, c)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "User Data Updated", "User": respU})

}

// TopUp godoc
// @Summary Top up account balance
// @Description must be authenticated as a donor user to top up account balance
// @Tags User
// @Accept  json
// @Produce  json
// @Param   Authorization  header  string  true  "Authentication token"  default()
// @Param topup body models.TopUpReq true "Top up request data"
// @Success 200 {object} map[string]interface{} "message: string, New Balance: float64"
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/user/topup [post]
func (h *UserHandler) TopUp(c echo.Context) error {
	cred := helper.GetCredential(c)
	if cred.Role != "donor" {
		return helper.ParseError(helper.ErrDonorUser, c)
	}
	var GetD models.TopUpReq
	err := c.Bind(&GetD)
	if err != nil {
		return helper.ParseError(helper.ErrBindJSON, c)
	}

	//validate deposit
	if GetD.Deposit <= 0 {
		return helper.ParseError(helper.ErrParam, c)
	}

	respU, err := h.UserGRPC.TopUp(
		context.TODO(),
		&pb.TopUpReq{Amount: GetD.Deposit, UserId: uint64(cred.UserID)},
	)
	if err != nil {
		return helper.ParseErrorGRPC(err, c)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Top Up success", "New Balance": respU.Balance})

}
