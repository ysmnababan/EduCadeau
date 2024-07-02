package router

import (
	"notification_service/internal/controller"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	// Route to send email
	e.POST("/send-email", controller.SendEmail)
}
