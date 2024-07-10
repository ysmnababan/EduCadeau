package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payment struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RegistryID    primitive.ObjectID `json:"registry_id" bson:"registry_id"`
	PaymentDate   string             `json:"payment_date" bson:"payment_date"`
	PaymentMethod string             `json:"payment_method" bson:"payment_method"`
	PaymentAmount float64            `json:"payment_amount" bson:"payment_amount"`
}

type PaymentReq struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RegistryID    primitive.ObjectID `json:"registry_id" bson:"registry_id"`
	PaymentDate   string             `json:"payment_date" bson:"payment_date"`
	PaymentMethod string             `json:"payment_method" bson:"payment_method"`
	PaymentAmount float64            `json:"payment_amount" bson:"payment_amount"`
}

type XenditInvoiceRequest struct {
	ExternalID  string  `json:"external_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

type XenditInvoiceResponse struct {
	ID         string `json:"id"`
	InvoiceURL string `json:"invoice_url"`
}
