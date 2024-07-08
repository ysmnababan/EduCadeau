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
	GetAllRegistries(filter string) ([]models.Registry, error)
	GetRegistryID(registry_id string, donor_id uint64) (models.Registry, error)
	Donate(in *models.Registry) error
	// FullDonate(in *models.Registry) (models.Registry, error)
	// PartialDonate(in *models.Registry) (models.Registry, error)
	DeleteRegistry(registry_id string, donor_id uint64) (string, error)
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
			return false, helper.ErrNoData
		}
		helper.Logging(nil).Error(err)
		return false, err
	}
	return true, nil
}

func (r *Repo) GetAllRegistries(filter string) ([]models.Registry, error) {
	var registries []models.Registry
	var cursor *mongo.Cursor
	var err error
	if filter == "" {
		cursor, err = r.DB.Collection("registries").Find(context.TODO(), bson.D{{}})
	} else {
		cursor, err = r.DB.Collection("registries").Find(context.TODO(), bson.M{"status": filter})
	}
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

	log.Printf("CREATE SUCCESS: %v\n\n", res)
	return nil
}

// func (r *Repo) PartialDonate(in *models.Registry) (models.Registry, error) {

// 	return models.Registry{}, nil
// }

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

	return fmt.Sprintf("REGISTRY WITH ID:%d IS DELETED SUCCESFULLY", r_id), nil
}
