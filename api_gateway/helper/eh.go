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
	ErrDonorUser           = errors.New("unauthorized, donor privilege only")
	ErrRecipientUser       = errors.New("unauthorized, recipient privilege only")
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

	case errors.Is(err, ErrDonorUser):
		status = http.StatusUnauthorized

	case errors.Is(err, ErrRecipientUser):
		status = http.StatusUnauthorized

	case errors.Is(err, ErrNoUpdate):
		status = http.StatusBadRequest

	default:
		status = http.StatusInternalServerError

	}

	return ctx.JSON(status, map[string]interface{}{"message": err.Error()})
}

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
