package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DonationDetail struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DonationID    primitive.ObjectID `json:"donation_id,omitempty" bson:"donation_id,omitempty"`
	Description   string             `json:"description,omitempty" bson:"description,omitempty"`
	DonationType  string             `json:"donation_type,omitempty" bson:"donation_type,omitempty"`
	Tag           string             `json:"tag,omitempty" bson:"tag,omitempty"`
	SenderAddress string             `json:"sender_address,omitempty" bson:"sender_address,omitempty"`
	RelatedLink   string             `json:"related_link,omitempty" bson:"related_link,omitempty"`
	Notes         string             `json:"notes,omitempty" bson:"notes,omitempty"`
}
