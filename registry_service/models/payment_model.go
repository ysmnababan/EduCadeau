package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payment struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RegistryID    primitive.ObjectID `json:"registry_id" bson:"registry_id"`
	PaymentDate   string             `json:"payment_date" bson:"payment_date"`
	PaymentMethod string             `json:"payment_method" bson:"payment_method"`
	PaymentAmount float64            `json:"payment_amount" bson:"payment_amount"`
	InvoiceLink   string             `json:"invoice_link" bson:"invoice_link"`
}


