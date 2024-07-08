package handler

import (
	"encoding/json"
	"log"
	"os"
	"reflect"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendToMail(body interface{}, subject string) {

	bodyValue := reflect.ValueOf(body)
	if bodyValue.Kind() == reflect.Ptr {
		bodyValue = bodyValue.Elem()
	}
	bodyBytes, err := json.Marshal(bodyValue.Interface())
	if err != nil {
		log.Println("Error marshalling body:", err)
		return
	}
	bodyStr := string(bodyBytes)

	// Create email message
	from := mail.NewEmail("educadeu_admin", "educadeu.service@gmail.com") // ubah jadi const
	to := mail.NewEmail("Recipient", "andhika.favian18@gmail.com")        // ubah jadi const
	message := mail.NewSingleEmail(from, subject, to, bodyStr, bodyStr)
	message.SetReplyTo(mail.NewEmail("educadeu_customer_service", "educadeu.service@gmail.com"))

	// Get SendGrid API key from environment variable
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		// log error
		return
	}

	// Send email
	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		// log error
		return
	}

	log.Println("SendGrid API LOGS:", response.StatusCode, response.Body)

}
