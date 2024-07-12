package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"notification_service/helper"
	"reflect"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Override default client to accept TLS 1.2 test host
func NewSendClient(key string, host string) *sendgrid.Client {
	request := sendgrid.GetRequest(key, "/v3/mail/send", host)
	request.Method = "POST"
	return &sendgrid.Client{Request: request}
}

func SendToMailTwilio(body interface{}, subject string) {
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
	htmlContent := fmt.Sprintf("<strong>%v</strong>", body)

	// Create email message
	from := mail.NewEmail("educadeu_admin", "educadeu.service@gmail.com") // ubah jadi const
	to := mail.NewEmail("Recipient", "educadeu.service@gmail.com")              // ubah jadi const
	message := mail.NewSingleEmail(from, subject, to, bodyStr, htmlContent)
	message.SetReplyTo(mail.NewEmail("educadeu_customer_service", "educadeu.service@gmail.com"))

	// Send email
	// Use TLS 1.2+ endpoint as host
	client := NewSendClient(helper.SENDGRID_API_KEY, "https://tlsv12.api.sendgrid.com")
	// client := sendgrid.NewSendClient(helper.SENDGRID_API_KEY)
	response, err := client.Send(message)
	if err != nil {
		log.Println("error sending message: ", err)
	} else {
		fmt.Println("status code", response.StatusCode)
		fmt.Println("response body", response.Body)
		fmt.Println("headers", response.Headers)
	}

}
func SendToMail(body interface{}, subject string) {
	// Ensure the interface is translated to string
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
	htmlContent := fmt.Sprintf("<strong>%v</strong>", body)

	// Create email message
	from := mail.NewEmail("educadeu_admin", "educadeu.service@gmail.com")
	to := mail.NewEmail("Recipient", "olansosmed@gmail.com")
	message := mail.NewSingleEmail(from, subject, to, bodyStr, htmlContent)
	message.SetReplyTo(mail.NewEmail("educadeu_customer_service", "olansosmed@gmail.com"))

	// Create a custom HTTP client with a timeout
	httpClient := &http.Client{
		Timeout: 120 * time.Second,
	}

	// Send the email
	reqBody, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshalling email message:", err)
		return
	}

	req, err := http.NewRequest("POST", "https://api.sendgrid.com/v3/mail/send", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer " +helper.SENDGRID_API_KEY)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("status code:", resp.StatusCode)
	// Read response body
	bodyResp := new(bytes.Buffer)
	bodyResp.ReadFrom(resp.Body)
	fmt.Println("response body:", bodyResp.String())
	fmt.Println("headers:", resp.Header)
}
