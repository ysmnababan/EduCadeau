package setup

import (
	"api_gateway/handler"
	"api_gateway/helper"
	"api_gateway/pb"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	User pb.UserToRestClient
}

func SetupRESTServer(e *echo.Echo, h *Handler) {
	e.Use(middleware.Recover())
	// using logger for each api
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			helper.Logging(c).Info("Calling an API")
			return next(c)
		}
	})

	userHandler := &handler.UserHandler{UserGRPC: h.User}
	e.POST("/login", userHandler.Login)
	e.POST("/register", userHandler.Register)

	protected := e.Group("")
	{
		protected.GET("/users", userHandler.GetAllUser)
		protected.GET("/user", userHandler.GetUserDetail)
		protected.PUT("/user/top-up", userHandler.TopUp)
		protected.PUT("/user", userHandler.EditUser)
	}
}
