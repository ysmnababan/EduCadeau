package email

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// EmailRequest defines the structure of the request body for sending an email
type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// SendEmail Funcion
func SendEmail(c echo.Context, request *EmailRequest) error {
	// Create email message
	from := mail.NewEmail("KUE_APEM", "educadeu.service@gmail.com")
	to := mail.NewEmail("Recipient", request.To)
	message := mail.NewSingleEmail(from, request.Subject, to, request.Body, request.Body)
	message.SetReplyTo(mail.NewEmail("KUE_APEM", "educadeu.service@gmail.com"))

	// Get SendGrid API key from environment variable
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "SendGrid API key not set"})
	}
	fmt.Println("API KEYYY:", apiKey)

	// Send email
	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Printf("SendGrid API error: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send email"})
	}

	// Log SendGrid response
	log.Printf("SendGrid API response: StatusCode=%d, Body=%s\n", response.StatusCode, response.Body)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Email sent successfully",
		"status":  response.StatusCode,
	})
}
