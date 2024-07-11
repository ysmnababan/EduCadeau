package handler

import (
	"encoding/json"
	"fmt"
	"log"

	"notification_service/helper"
	"reflect"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendToMail(body interface{}, subject string) {
	// Make sure the interface is translated to string
	bodyValue := reflect.ValueOf(body)
	if bodyValue.Kind() == reflect.Ptr {
		bodyValue = bodyValue.Elem()
	}
	log.Println("inside mail message", bodyValue.Interface())

	bodyBytes, err := json.Marshal(bodyValue.Interface())
	if err != nil {
		log.Println("Error marshalling body:", err)
		return
	}
	bodyStr := string(bodyBytes)

	// Create email message
	from := mail.NewEmail("educadeu_admin", "educadeu.service@gmail.com") // ubah jadi const
	to := mail.NewEmail("Recipient", "educadeu.service@gmail.com")        // ubah jadi const
	message := mail.NewSingleEmail(from, subject, to, bodyStr, bodyStr)
	message.SetReplyTo(mail.NewEmail("educadeu_customer_service", "educadeu.service@gmail.com"))

	// Send email

	client := sendgrid.NewSendClient(helper.SENDGRID_API_KEY)
	response, err := client.Send(message)
	if err != nil {
		log.Println("error sending message: ", err)
	} else {
		fmt.Println("status code", response.StatusCode)
		fmt.Println("response body", response.Body)
		fmt.Println("headers", response.Headers)
	}

}
