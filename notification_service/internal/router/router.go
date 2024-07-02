// internal/router/router.go

package router

import (
	"notification_service/internal/controller"

	"github.com/labstack/echo/v4"
)

// InitRoutes initializes routes for the application
func InitRoutes(e *echo.Echo) {
	// Route to send email
	e.POST("/send-email", controller.SendEmail)
}
