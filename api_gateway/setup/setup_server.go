package setup

import (
	"api_gateway/handler"
	"api_gateway/helper"
	"api_gateway/pb"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.elastic.co/apm/module/apmechov4/v2"
)

type Handler struct {
	User pb.UserToRestClient
}

func SetupRESTServer(e *echo.Echo, h *Handler) {
	e.Use(middleware.Recover())
	e.Use(apmechov4.Middleware())

	// using logger for each api
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Print("\033[H\033[2J")
			fmt.Println("==================================")
			fmt.Println("            EDU CADEU")
			fmt.Println("==================================")

			// Wrap the response writer to capture the status code
			rr := &helper.ResponseRecorder{ResponseWriter: c.Response().Writer, Status: http.StatusOK}
			c.Response().Writer = rr

			err := next(c)

			// Log the request details including the status code
			helper.Logging(c).Info("request detail")
			return err
		}
	})

	userHandler := &handler.UserHandler{UserGRPC: h.User}
	e.POST("/login", userHandler.Login)
	e.POST("/register", userHandler.Register)

	protected := e.Group("")
	protected.Use(helper.Auth)
	{
		protected.GET("/users", userHandler.GetAllUser)
		protected.GET("/user", userHandler.GetUserDetail)
		protected.PUT("/user/top-up", userHandler.TopUp)
		protected.PUT("/user", userHandler.EditUser)
	}
}
