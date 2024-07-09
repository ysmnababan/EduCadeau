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

// func (r *Repo) isPaymentExist(filter any) (bool, error) {
// 	var result bson.M
// 	err := r.DB.Collection("payments").FindOne(context.TODO(), filter).Decode(&result)
// 	if err != nil {
// 		// transaction not found
// 		if err == mongo.ErrNoDocuments {
// 			return false, helper.ErrNoData
// 		}
// 		helper.Logging(nil).Error(err)
// 		return false, helper.ErrQuery
// 	}

// 	return true, nil
// }

func (r *Repo) isPaymentValid(registry_id primitive.ObjectID, donor_id uint) (bool, error) {
	var result bson.M
	err := r.DB.Collection("registries").FindOne(context.TODO(), bson.M{"_id": registry_id, "donor_id": donor_id}).Decode(&result)
	if err != nil {
		// registries not found
		if err == mongo.ErrNoDocuments {
			return false, helper.ErrNoData
		}
		helper.Logging(nil).Error(err)
		return false, helper.ErrQuery
	}

	err = r.DB.Collection("payments").FindOne(context.TODO(), bson.M{"registry_id": registry_id}).Decode(&result)
	if err != nil && err != mongo.ErrNoDocuments {
		helper.Logging(nil).Error(err)
		return false, helper.ErrQuery
	}
	if err == nil {
		helper.Logging(nil).Error("ERROR REPO: REGISTRY ALREADY EXISTS IN PAYMENT")
		return false, helper.ErrParam
	}
	return true, nil
}

func (r *Repo) GetAllPayments() ([]*models.Payment, error) {
	var payments []*models.Payment
	cursor, err := r.DB.Collection("payments").Find(context.TODO(), bson.D{{}})
	if err != nil {
		helper.Logging(nil).Error("ERROR REPO: ", err)
		return nil, helper.ErrQuery
	}

	for cursor.Next(context.TODO()) {
		var d models.Payment
		if err := cursor.Decode(&d); err != nil {
			helper.Logging(nil).Error("ERROR REPO: ", err)
			return nil, helper.ErrQuery
		}

		payments = append(payments, &d)
	}

	log.Printf("GET ALL DATA SUCCESS: %v\n\n", payments)
	return payments, nil

}

func (r *Repo) GetAllMyPayments(donor_id uint64) ([]*models.Payment, error) {
	cursor, err := r.DB.Collection("registries").Find(context.TODO(), bson.M{"donor_id": donor_id, "status": "settlement"})
	if err != nil {
		helper.Logging(nil).Error("ERROR REPO: ", err)
		return nil, helper.ErrQuery
	}
	log.Println("here")
	out := []*models.Payment{}
	for cursor.Next(context.TODO()) {
		var reg models.Registry
		if err := cursor.Decode(&reg); err != nil {
			helper.Logging(nil).Error("ERROR REPO: ", err)
			return nil, helper.ErrQuery
		}
		log.Println("here", reg)
		var p models.Payment
		err = r.DB.Collection("payments").FindOne(context.TODO(), bson.M{"registry_id": reg.ID}).Decode(&p)
		if err != nil {
			helper.Logging(nil).Error("ERROR REPO: ", err)
			continue
		}
		log.Println("here")

		out = append(out, &p)
	}

	log.Println(out)
	return out, nil
}

func (r *Repo) GetPayment(donor_id uint64, payment_id string) (*models.Payment, error) {
	var result models.Payment
	p_id, err := primitive.ObjectIDFromHex(payment_id)
	if err != nil {
		return nil, helper.ErrInvalidId
	}
	err = r.DB.Collection("payments").FindOne(context.TODO(), bson.M{"_id": p_id}).Decode(&result)
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
	err = r.DB.Collection("registries").FindOne(context.TODO(), bson.M{"donor_id": donor_id, "_id": result.RegistryID}).Decode(&res)
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

func (r *Repo) Pay(in *models.Payment, donor_id uint64) error {
	isValid, err := r.isPaymentValid(in.RegistryID, uint(donor_id))
	if err != nil || !isValid {
		return helper.ErrPaymentSettled
	}

	res, err := r.DB.Collection("payments").InsertOne(context.TODO(), *in)
	if err != nil {
		helper.Logging(nil).Error("CREATE PAYMENT: ", err)
		return helper.ErrQuery
	}
	in.ID = res.InsertedID.(primitive.ObjectID)

	log.Println("PAY A DONATION: ", res)
	update, err := r.DB.Collection("registries").UpdateOne(
		context.TODO(),
		bson.M{"_id": in.RegistryID},
		bson.M{"$set": bson.M{"status": "settlement"}},
	)
	if err != nil {
		helper.Logging(nil).Error("UPDATE STATUS PAYMENT: ", err)
		return helper.ErrQuery
	}
	log.Println("PAY A DONATION: ", update)

	return nil
}
