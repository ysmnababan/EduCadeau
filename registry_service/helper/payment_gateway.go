package helper

import "registry_service/pb/pbDonationRegistry"

func PaymentGateway(amount float64, detail *pbDonationRegistry.DonationResp) (string, error) {
	// write your code here
	return "invoicelink.com", nil
}
