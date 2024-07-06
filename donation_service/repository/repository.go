package repository

import (
	"donation_service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	DB *mongo.Database
}

type DonationRepo interface {
	GetAllDonations(filter string) ([]models.Donation, error)
	GetDonationDetail(donation_id string) (models.DonationDetailResp, error)
	CreateDonation(d *models.CreateDonationReq) (models.DonationDetailResp, error)
	EditDonation(d *models.EditDonationReq) (models.DonationDetailResp, error)
	DeleteDonation(donation_id string) (interface{}, error)
}

func (r *Repo) GetAllDonations(filter string) ([]models.Donation, error) {

	return nil, nil
}

func (r *Repo) GetDonationDetail(donation_id string) (models.DonationDetailResp, error) {

	return models.DonationDetailResp{}, nil
}

func (r *Repo) CreateDonation(d *models.CreateDonationReq) (models.DonationDetailResp, error) {

	return models.DonationDetailResp{}, nil
}

func (r *Repo) EditDonation(d *models.EditDonationReq) (models.DonationDetailResp, error) {

	return models.DonationDetailResp{}, nil
}

func (r *Repo) DeleteDonation(donation_id string) (interface{}, error) {

	return nil, nil
}
