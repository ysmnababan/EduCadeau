package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Registry struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DonationID primitive.ObjectID `json:"donation_id" bson:"donation_id"`
	DonorID    uint64             `json:"donor_id" bson:"donor_id"`
	Amount     float64            `json:"amount" bson:"amount"`
	Status     string             `json:"status" bson:"status"`
}

type CreateRegistryReq struct {
	DonationID string  `json:"donation_id,omitempty" bson:"donation_id,omitempty"`
	Amount     float64 `json:"amount" bson:"amount"`
	Filter     string  `json:"filter" bson:"filter"`
}
