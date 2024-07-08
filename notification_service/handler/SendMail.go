package handler

import (
	"log"
	"net/http"
	"os"

	"notification_service/models"

	"github.com/labstack/echo/v4"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(c echo.Context, Request struct{}) error {
	// Parse request body
	var request models.Request
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Create email message
	from := mail.NewEmail("educadeu_admin", "educadeu.service@gmail.com") // ubah jadi const
	to := mail.NewEmail("Recipient", request.To)                          // ubah jadi const
	message := mail.NewSingleEmail(from, request.Subject, to, request.Body, request.Body)
	message.SetReplyTo(mail.NewEmail("educadeu_customer_service", "educadeu.service@gmail.com"))

	// Get SendGrid API key from environment variable
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "SendGrid API key not set"})
	}

	// Send email
	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send email"})
	}

	log.Println("SendGrid API LOGS:", response.StatusCode, response.Body)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Email sent successfully",
		"status":  response.StatusCode,
	})
}
