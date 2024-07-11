package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"registry_service/pb/pbDonationRegistry"
)

const (
	XenditURL = "https://api.xendit.co/v2/invoices"
)

type XenditInvoiceRequest struct {
	DonationName  string  `json:"donation_name"`
	RecipientName string  `json:"recipient_name"`
	ExternalID    string  `json:"external_id"`
	Amount        float64 `json:"amount"`
	Description   string  `json:"description"`
	Status        string  `json:"status"`
}

type XenditInvoiceResponse struct {
	ID         string `json:"id"`
	InvoiceURL string `json:"invoice_url"`
}

func PaymentGateway(amount float64, detail *pbDonationRegistry.DonationResp) (string, error) {
	// Prepare the request payload
	requestPayload := XenditInvoiceRequest{
		DonationName:  detail.DonationName,
		RecipientName: detail.RecipientName,
		ExternalID:    fmt.Sprintf("donation_%d", detail.RecipientId),
		Amount:        amount,
		Description:   detail.Description,
		Status:        fmt.Sprintf("status:%v", detail.Status),
	}

	payloadBytes, err := json.Marshal(requestPayload)
	if err != nil {
		log.Fatalf("Error marshalling request payload: %v", err)
	}

	req, err := http.NewRequest("POST", XenditURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(XENDIT_SECRET_KEY, "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body for detailed error message
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error response from Xendit: %v\nResponse body: %s", resp.Status, string(body))
	}


	var invoiceResponse XenditInvoiceResponse
	err = json.Unmarshal(body, &invoiceResponse)
	if err != nil {
		log.Fatalf("Error decoding response body: %v", err)
	}
	fmt.Println(invoiceResponse.ID)
	fmt.Printf("Invoice created: %s\n", invoiceResponse.InvoiceURL)

	return invoiceResponse.InvoiceURL, nil

}
