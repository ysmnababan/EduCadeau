package helper

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNoData              = errors.New("no data in result set")
	ErrNoUser              = errors.New("no user exist")
	ErrQuery               = errors.New("query execution failed")
	ErrInvalidId           = errors.New("invalid id")
	ErrUserExists          = errors.New("user already exist")
	ErrNoUpdate            = errors.New("data already exists")
	ErrBindJSON            = errors.New("unable to bind json")
	ErrParam               = errors.New("error or missing parameter")
	ErrCredential          = errors.New("password or email doesn't match")
	ErrGeneratedPwd        = errors.New("error generating password hash")
	ErrMustAdmin           = errors.New("unauthorized, admin privilege only")
	ErrOnlyUser            = errors.New("unauthorized, user privilege only")
	ErrUnsufficientBalance = errors.New("no sufficient fund")
)

func ParseError(err error, ctx echo.Context) error {
	status := http.StatusOK

	switch {
	case errors.Is(err, ErrQuery):
		fallthrough
	case errors.Is(err, ErrGeneratedPwd):
		fallthrough
	case errors.Is(err, ErrNoUser):
		status = http.StatusNotFound

	case errors.Is(err, ErrNoData):
		status = http.StatusNotFound

	case errors.Is(err, ErrUnsufficientBalance):
		status = http.StatusBadRequest

	case errors.Is(err, ErrParam):
		status = http.StatusBadRequest

	case errors.Is(err, ErrBindJSON):
		status = http.StatusBadRequest

	case errors.Is(err, ErrInvalidId):
		status = http.StatusBadRequest

	case errors.Is(err, ErrCredential):
		status = http.StatusBadRequest

	case errors.Is(err, ErrUserExists):
		status = http.StatusBadRequest

	case errors.Is(err, ErrMustAdmin):
		status = http.StatusUnauthorized

	case errors.Is(err, ErrOnlyUser):
		status = http.StatusUnauthorized

	case errors.Is(err, ErrNoUpdate):
		status = http.StatusBadRequest

	default:
		status = http.StatusInternalServerError

	}

	return ctx.JSON(status, map[string]interface{}{"message": err.Error()})
}

// func ParseErrorGRPC(err error) error {
// 	Logging(nil).Error("ERR: ", err)
// 	code := codes.OK
// 	message := ""
// 	switch {
// 	case errors.Is(err, ErrNoUser):
// 		code = codes.NotFound
// 		message = "No User found"
// 	case errors.Is(err, ErrNoData):
// 		code = codes.NotFound
// 		message = "No data found"
// 	case errors.Is(err, ErrUnsufficientBalance):
// 		code = codes.InvalidArgument
// 		message = "unsufficient balance"
// 	case errors.Is(err, ErrParam):
// 		code = codes.InvalidArgument
// 		message = "error or missing param"
// 	case errors.Is(err, ErrBindJSON):
// 		code = codes.InvalidArgument
// 		message = "Bad request"
// 	case errors.Is(err, ErrInvalidId):
// 		code = codes.InvalidArgument
// 		message = "Invalid ID"
// 	case errors.Is(err, ErrCredential):
// 		code = codes.InvalidArgument
// 		message = "email or password missmatch"
// 	case errors.Is(err, ErrUserExists):
// 		code = codes.AlreadyExists
// 		message = "User Already Exists"
// 	case errors.Is(err, ErrMustAdmin):
// 		code = codes.PermissionDenied
// 		message = "Admin privilege only"
// 	case errors.Is(err, ErrOnlyUser):
// 		code = codes.PermissionDenied
// 		message = "User privilege only"
// 	case errors.Is(err, ErrNoUpdate):
// 		code = codes.AlreadyExists
// 		message = "Data is the same"
// 	case errors.Is(err, ErrQuery):
// 		fallthrough
// 	case errors.Is(err, ErrGeneratedPwd):
// 		fallthrough
// 	default:
// 		code = codes.Internal
// 		message = "Unknown error:" + err.Error()
// 	}

// 	// log.Println(map[string]interface{}{"message": message, "status": status})
// 	return status.Errorf(code, message)
// }

func ParseErrorGRPC(err error, ctx echo.Context) error {
	stat := http.StatusOK
	message := ""
	if st, ok := status.FromError(err); ok {
		message = st.Message()

		switch st.Code() {
		case codes.NotFound:
			stat = http.StatusNotFound
		case codes.InvalidArgument:
			stat = http.StatusBadRequest
		case codes.AlreadyExists:
			stat = http.StatusBadRequest
		case codes.PermissionDenied:
			stat = http.StatusUnauthorized
		case codes.Internal:
			fallthrough
		default:
			log.Println(st.Code())
			stat = http.StatusInternalServerError
			message = "Unknown error:" + err.Error()
		}
	} else {
		fmt.Printf("not able to parse error returned %v", err)
		message = "Internal error parsing grpc"
	}
	log.Println("ERROR=> ", err)
	return ctx.JSON(stat, map[string]interface{}{"message": message})
}
