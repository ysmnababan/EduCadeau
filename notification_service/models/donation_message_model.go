package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DonationDetailResp struct {
	ID                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RecipientID       uint               `json:"recipient_id,omitempty" bson:"recipient_id,omitempty"`
	RecipientName     string
	DonationName      string  `json:"donation_name" bson:"donation_name"`
	CreatedAt         string  `json:"created_at" bson:"created_at"`
	Status            string  `json:"status" bson:"status"`
	TargetAmount      float64 `json:"target_amount" bson:"target_amount"`
	AmountCollected   float64 `json:"amount_collected" bson:"amount_collected"`
	MiscellaneousCost float64 `json:"miscellaneous_cost,omitempty" bson:"miscellaneous_cost,omitempty"`
	Description       string  `json:"description,omitempty" bson:"description,omitempty"`
	DonationType      string  `json:"donation_type,omitempty" bson:"donation_type,omitempty"`
	Tag               string  `json:"tag,omitempty" bson:"tag,omitempty"`
	SenderAddress     string  `json:"sender_address,omitempty" bson:"sender_address,omitempty"`
	RelatedLink       string  `json:"related_link,omitempty" bson:"related_link,omitempty"`
	Notes             string  `json:"notes,omitempty" bson:"notes,omitempty"`
}
