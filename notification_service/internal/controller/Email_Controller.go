// internal/controller/email_controller.go

package controller

import (
	"net/http"

	"notification_service/internal/email"

	"github.com/labstack/echo/v4"
)

// SendEmail handles sending an email
func SendEmail(c echo.Context) error {
	// Parse request body
	var request email.EmailRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Call the SendEmail function from the email package
	return email.SendEmail(c, &request)
}
