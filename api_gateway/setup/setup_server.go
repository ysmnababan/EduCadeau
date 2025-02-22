package setup

import (
	"api_gateway/handler"
	"api_gateway/helper"
	"api_gateway/pb"
	"api_gateway/pb/donation_rest"
	"api_gateway/pb/pbRegistryRest"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.elastic.co/apm/module/apmechov4/v2"
)

type Handler struct {
	User     pb.UserToRestClient
	Donation donation_rest.DonationRestClient
	Registry pbRegistryRest.RegistryRestClient
}

// @title Edu Cadeu API
// @version 1.0
// @description This is API documentation for Edu Cadeu
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @host api-gateway-753cnjdw3a-et.a.run.app
func SetupRESTServer(e *echo.Echo, h *Handler) {
	e.Use(middleware.Recover())
	e.Use(apmechov4.Middleware())

	// using logger for each api
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// fmt.Print("\033[H\033[2J")
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
	donationHandler := &handler.DonationHandler{DonationGRPC: h.Donation}
	registryHandler := &handler.RegistryHandler{RegistryGRPC: h.Registry}

	e.POST("/login", userHandler.Login)
	e.POST("/register", userHandler.Register)

	protected := e.Group("")
	protected.Use(helper.Auth)
	{

		// for user
		protected.GET("/users", userHandler.GetAllUser)
		protected.GET("/user", userHandler.GetUserDetail)
		protected.PUT("/user/top-up", userHandler.TopUp)
		protected.PUT("/user", userHandler.EditUser)

		// for donation
		protected.GET("/donations", donationHandler.GetAllDonations)
		protected.GET("/donation/:id", donationHandler.GetDonationDetail)
		protected.POST("/donation", donationHandler.CreateDonation)
		protected.PUT("/donation/:id", donationHandler.EditDonation)
		protected.DELETE("/donation/:id", donationHandler.DeleteDonation)

		// for registry
		protected.GET("/donated", registryHandler.GetAllRegistries)
		protected.GET("/donated/:id", registryHandler.DetailOfRegistry)
		protected.POST("/donate", registryHandler.Donate)
		protected.DELETE("/donated/:id", registryHandler.DeleteRegistry)

		// for payments
		protected.GET("/payments", registryHandler.GetAllPayments)
		protected.GET("/payment/:id", registryHandler.GetPaymentID)
		protected.POST("/payment/:id", registryHandler.PayDonation)
	}
}
