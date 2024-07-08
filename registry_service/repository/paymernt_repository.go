package repository

import (
	"context"
	"log"
	"registry_service/helper"
	"registry_service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repo) isPaymentExist(payment_id primitive.ObjectID) (bool, error) {
	var result bson.M
	err := r.DB.Collection("payments").FindOne(context.TODO(), bson.M{"_id": payment_id}).Decode(&result)
	if err != nil {
		// transaction not found
		if err == mongo.ErrNoDocuments {
			return false, helper.ErrNoData
		}
		helper.Logging(nil).Error(err)
		return false, err
	}

	return true, nil
}

func (r *Repo) isPaymentDuplicate(donation_id primitive.ObjectID, donor_id uint) (bool, error) {
	var result bson.M

	err := r.DB.Collection("registries").FindOne(context.TODO(), bson.M{"_id": donation_id, "donor_id": donor_id}).Decode(&result)

	if err != nil {
		// transaction not found
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		helper.Logging(nil).Error(err)
		return false, err
	}
	return true, helper.ErrUserExists
}

func (r *Repo) GetAllPayments() ([]*models.Payment, error) {
	var payments []*models.Payment
	cursor, err := r.DB.Collection("payments").Find(context.TODO(), bson.D{{}})
	if err != nil {
		helper.Logging(nil).Error("ERROR REPO: ", err)
		return nil, helper.ErrQuery
	}

	for cursor.Next(context.TODO()) {
		var d *models.Payment
		if err := cursor.Decode(d); err != nil {
			helper.Logging(nil).Error("ERROR REPO: ", err)
			return nil, helper.ErrQuery
		}

		payments = append(payments, d)
	}

	log.Printf("GET ALL DATA SUCCESS: %v\n\n", payments)
	return payments, nil

}

func (r *Repo) GetAllMyPayments(donor_id uint64) ([]*models.Payment, error) {
	cursor, err := r.DB.Collection("registries").Find(context.TODO(), bson.M{"donor_id": donor_id})
	if err != nil {
		helper.Logging(nil).Error("ERROR REPO: ", err)
		return nil, helper.ErrQuery
	}

	out := []*models.Payment{}
	for cursor.Next(context.TODO()) {
		var reg models.Registry
		if err := cursor.Decode(&reg); err != nil {
			helper.Logging(nil).Error("ERROR REPO: ", err)
			return nil, helper.ErrQuery
		}
		var p models.Payment
		err = r.DB.Collection("payments").FindOne(context.TODO(), bson.M{"registry_id": reg.ID.Hex()}).Decode(&p)
		if err != nil {
			helper.Logging(nil).Error("ERROR REPO: ", err)
			continue
		}

		out = append(out, &p)
	}

	log.Println(out)
	return out, nil
}

func (r *Repo) GetPayment(donor_id uint64, payment_id string) (*models.Payment, error) {
	var result models.Payment
	err := r.DB.Collection("payments").FindOne(context.TODO(), bson.M{"_id": payment_id}).Decode(&result)
	if err != nil {
		// payment not found
		if err == mongo.ErrNoDocuments {
			helper.Logging(nil).Error("REPO: NO PAYMENT")
			return nil, helper.ErrNoData
		}
		helper.Logging(nil).Error(err)
		return nil, err
	}

	log.Println("here")
	var res bson.M
	err = r.DB.Collection("registries").FindOne(context.TODO(), bson.M{"donor_id": donor_id, "_id": result.RegistryID}).Decode(res)
	if err != nil {
		// payment not found
		if err == mongo.ErrNoDocuments {
			helper.Logging(nil).Error("REPO: NO REGISTRIES")
			return nil, helper.ErrNoData
		}
		helper.Logging(nil).Error(err)
		return nil, err
	}

	return &result, nil
}

func (r *Repo) Pay(in *models.Payment) error {

	return nil
}
