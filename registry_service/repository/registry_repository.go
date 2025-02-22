package repository

import (
	"context"
	"fmt"
	"log"
	"registry_service/helper"
	"registry_service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	DB *mongo.Database
}

type RegistryRepo interface {
	GetAllRegistries(filter string, donor_id uint) ([]models.Registry, error)
	GetRegistryID(registry_id string, donor_id uint64) (models.Registry, error)
	Donate(in *models.Registry) error
	DeleteRegistry(registry_id string, donor_id uint64) (string, error)

	GetAllPayments() ([]*models.Payment, error)
	GetAllMyPayments(donor_id uint64) ([]*models.Payment, error)
	GetPayment(donor_id uint64, payment_id string) (*models.Payment, error)
	Pay(in *models.Payment, donor_id uint64) error
}

func (r *Repo) isRegistryExist(registry_id primitive.ObjectID, donor_id uint) (bool, error) {
	var result bson.M
	var err error
	if donor_id == 0 {
		err = r.DB.Collection("registries").FindOne(context.TODO(), bson.M{"_id": registry_id}).Decode(&result)
	} else {
		err = r.DB.Collection("registries").FindOne(context.TODO(), bson.M{"_id": registry_id, "donor_id": donor_id}).Decode(&result)
	}

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

func (r *Repo) isRegistryDuplicate(donation_id primitive.ObjectID, donor_id uint) (bool, error) {
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

func (r *Repo) GetAllRegistries(filter string, donor_id uint) ([]models.Registry, error) {
	var registries []models.Registry
	var cursor *mongo.Cursor
	var err error

	var qFilter interface{}
	if filter == "" && donor_id == 0 {
		qFilter = bson.D{{}}
	} else if filter == "" && donor_id != 0 {
		qFilter = bson.M{"donor_id": donor_id}
	} else if filter != "" && donor_id == 0 {
		qFilter = bson.M{"status": filter}
	} else if filter != "" && donor_id != 0 {
		qFilter = bson.M{"status": filter, "donor_id": donor_id}
	}

	cursor, err = r.DB.Collection("registries").Find(context.TODO(), qFilter)

	if err != nil {
		helper.Logging(nil).Error("ERROR REPO: ", err)
		return nil, helper.ErrQuery
	}

	for cursor.Next(context.TODO()) {
		var d models.Registry
		if err := cursor.Decode(&d); err != nil {
			helper.Logging(nil).Error("ERROR REPO: ", err)
			return nil, helper.ErrQuery
		}

		registries = append(registries, d)
	}

	log.Printf("GET ALL DATA SUCCESS: %v\n\n", registries)
	return registries, nil
}

func (r *Repo) GetRegistryID(registry_id string, donor_id uint64) (models.Registry, error) {
	var rOut models.Registry

	r_id, err := primitive.ObjectIDFromHex(registry_id)
	if err != nil {
		helper.Logging(nil).Error("ERR REPO: ", err)
		return models.Registry{}, helper.ErrInvalidId
	}

	isRegistryExist, err := r.isRegistryExist(r_id, uint(donor_id))
	if err != nil {
		return models.Registry{}, err
	}

	if !isRegistryExist {
		return models.Registry{}, helper.ErrNoData
	}

	r.DB.Collection("registries").FindOne(context.TODO(), bson.M{"_id": r_id}).Decode(&rOut)
	log.Printf("GET DATA SUCCESS: %v\n\n", rOut)

	return rOut, nil
}

func (r *Repo) Donate(in *models.Registry) error {
	isDuplicate, err := r.isRegistryDuplicate(in.DonationID, uint(in.DonorID))
	if err != nil {
		return err
	}

	if isDuplicate {
		return helper.ErrUserExists
	}

	in.Status = "pending"
	res, err := r.DB.Collection("registries").InsertOne(context.TODO(), *in)
	if err != nil {
		helper.Logging(nil).Error("ERROR REPO: ", err)
		return err
	}

	in.ID = res.InsertedID.(primitive.ObjectID)

	log.Printf("CREATE SUCCESS: %v\n\n", res)
	return nil
}


func (r *Repo) DeleteRegistry(registry_id string, donor_id uint64) (string, error) {
	r_id, err := primitive.ObjectIDFromHex(registry_id)
	if err != nil {
		helper.Logging(nil).Error("ERR REPO: ", err)
		return "", helper.ErrInvalidId
	}

	isRegistryExist, err := r.isRegistryExist(r_id, uint(donor_id))
	if err != nil {
		return "", err
	}

	if !isRegistryExist {
		return "", helper.ErrNoData
	}

	res, err := r.DB.Collection("registries").DeleteOne(context.TODO(), bson.M{"_id": r_id})
	if err != nil {
		helper.Logging(nil).Error("ERROR REPO: ", err)
		return "", err
	}

	log.Printf("DELETE SUCCESS: %v\n\n", res)

	return fmt.Sprintf("REGISTRY WITH ID:%v IS DELETED SUCCESFULLY", r_id.Hex()), nil
}
