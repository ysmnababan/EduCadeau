package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Donation struct {
	ID                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RecipientID       uint               `json:"recipient_id,omitempty" bson:"recipient_id,omitempty"`
	DonationName      string             `json:"donation_name,omitempty" bson:"donation_name,omitempty"`
	CreatedAt         string             `json:"created_at,omitempty" bson:"created_at,omitempty"`
	Status            string             `json:"status,omitempty" bson:"status,omitempty"`
	TargetAmount      float64            `json:"target_amount,omitempty" bson:"target_amount,omitempty"`
	AmountCollected   float64            `json:"amount_collected,omitempty" bson:"amount_collected,omitempty"`
	MiscellaneousCost float64            `json:"miscellaneous_cost,omitempty" bson:"miscellaneous_cost,omitempty"`
}
