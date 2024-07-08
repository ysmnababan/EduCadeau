package repository

import (
	"registry_service/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	DB *mongo.Database
}

type RegistryRepo interface {
	GetAllRegistries(filter string) ([]models.Registry, error)
	GetRegistryID(registry_id primitive.ObjectID, donor_id uint64) (models.Registry, error)
	FullDonate(in *models.Registry) (models.Registry, error)
	PartialDonate(in *models.Registry) (models.Registry, error)
	DeleteRegistry(registry_id primitive.ObjectID, donor_id uint64) (string, error)
}

func (r *Repo) GetAllRegistries(filter string) ([]models.Registry, error) {

	return nil, nil
}

func (r *Repo) GetRegistryID(registry_id primitive.ObjectID, donor_id uint64) (models.Registry, error) {

	return models.Registry{}, nil
}

func (r *Repo) FullDonate(in *models.Registry) (models.Registry, error) {

	return models.Registry{}, nil
}

func (r *Repo) PartialDonate(in *models.Registry) (models.Registry, error) {

	return models.Registry{}, nil
}

func (r *Repo) DeleteRegistry(registry_id primitive.ObjectID, donor_id uint64) (string, error) {

	return "", nil
}
