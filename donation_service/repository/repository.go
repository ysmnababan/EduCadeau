package repository

import (
	"context"
	"donation_service/helper"
	"donation_service/models"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	DB *mongo.Database
}

type DonationRepo interface {
	GetAllDonations(filter string) ([]models.Donation, error)
	GetDonationDetail(donation_id string, recipient_id uint) (models.DonationDetailResp, error)
	CreateDonation(dreq *models.CreateDonationReq) (models.DonationDetailResp, error)
	EditDonation(dreq *models.EditDonationReq) (models.DonationDetailResp, error)
	DeleteDonation(donation_id string, recipient_id uint) (string, error)
}

func (r *Repo) isDonationExist(donation_id primitive.ObjectID, recipient_id uint) (bool, error) {
	var result bson.M
	var err error
	if recipient_id == 0 {
		err = r.DB.Collection("donations").FindOne(context.TODO(), bson.M{"_id": donation_id}).Decode(&result)
	} else {
		err = r.DB.Collection("donations").FindOne(context.TODO(), bson.M{"_id": donation_id, "recipient_id": recipient_id}).Decode(&result)
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

func (r *Repo) GetAllDonations(filter string) ([]models.Donation, error) {
	var donations []models.Donation
	cursor, err := r.DB.Collection("donations").Find(context.TODO(), bson.M{"status": filter})
	if err != nil {
		helper.Logging(nil).Error("ERROR REPO: ", err)
		return nil, helper.ErrQuery
	}

	for cursor.Next(context.TODO()) {
		var d models.Donation
		if err := cursor.Decode(&d); err != nil {
			helper.Logging(nil).Error("ERROR REPO: ", err)
			return nil, helper.ErrQuery
		}

		donations = append(donations, d)
	}

	return donations, nil
}

func (r *Repo) GetDonationDetail(donation_id string, recipient_id uint) (models.DonationDetailResp, error) {
	var d models.Donation
	var dd models.DonationDetail

	d_id, err := primitive.ObjectIDFromHex(donation_id)
	if err != nil {
		helper.Logging(nil).Error("ERR REPO: ", err)
		return models.DonationDetailResp{}, helper.ErrInvalidId
	}

	isDonationExist, err := r.isDonationExist(d_id, recipient_id)
	if err != nil {
		return models.DonationDetailResp{}, err
	}

	if !isDonationExist {
		return models.DonationDetailResp{}, helper.ErrNoData
	}

	r.DB.Collection("donations").FindOne(context.TODO(), bson.M{"_id": d_id}).Decode(&d)
	r.DB.Collection("donation_details").FindOne(context.TODO(), bson.M{"donation_id": d_id}).Decode(&dd)

	resp := models.DonationDetailResp{
		ID:                d.ID,
		RecipientID:       d.RecipientID,
		DonationName:      d.DonationName,
		CreatedAt:         d.CreatedAt,
		Status:            d.Status,
		TargetAmount:      d.TargetAmount,
		AmountCollected:   d.AmountCollected,
		MiscellaneousCost: d.MiscellaneousCost,
		Description:       dd.Description,
		DonationType:      dd.DonationType,
		Tag:               dd.Tag,
		SenderAddress:     dd.SenderAddress,
		RelatedLink:       dd.RelatedLink,
		Notes:             dd.Notes,
	}

	log.Println("donation_detail_resp: ", resp)
	return resp, nil
}

func (r *Repo) CreateDonation(dreq *models.CreateDonationReq) (models.DonationDetailResp, error) {
	d := models.Donation{
		RecipientID:       dreq.RecipientID,
		DonationName:      dreq.DonationName,
		AmountCollected:   dreq.AmountCollected,
		MiscellaneousCost: dreq.MiscellaneousCost,
	}
	d.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	d.Status = "no donor"
	d.AmountCollected = 0

	res, err := r.DB.Collection("donations").InsertOne(
		context.TODO(),
		d,
	)
	if err != nil {
		helper.Logging(nil).Error("ERR REPO: ", err)
		return models.DonationDetailResp{}, helper.ErrQuery
	}

	dd := models.DonationDetail{
		DonationID:    res.InsertedID.(primitive.ObjectID),
		Description:   dreq.Description,
		DonationType:  dreq.DonationType,
		Tag:           dreq.Tag,
		SenderAddress: dreq.SenderAddress,
		RelatedLink:   dreq.RelatedLink,
		Notes:         dreq.Notes,
	}

	res, err = r.DB.Collection("donation_details").InsertOne(
		context.TODO(),
		dd,
	)
	if err != nil {
		helper.Logging(nil).Error("ERR REPO: ", err)
		return models.DonationDetailResp{}, helper.ErrQuery
	}

	log.Println("Donation Detail ID: ", res.InsertedID)

	resp := models.DonationDetailResp{
		ID:                d.ID,
		RecipientID:       d.RecipientID,
		DonationName:      d.DonationName,
		CreatedAt:         d.CreatedAt,
		Status:            d.Status,
		TargetAmount:      d.TargetAmount,
		AmountCollected:   d.AmountCollected,
		MiscellaneousCost: d.MiscellaneousCost,
		Description:       dd.Description,
		DonationType:      dd.DonationType,
		Tag:               dd.Tag,
		SenderAddress:     dd.SenderAddress,
		RelatedLink:       dd.RelatedLink,
		Notes:             dd.Notes,
	}

	log.Println("donation_detail_resp: ", resp)
	return resp, nil
}

func (r *Repo) EditDonation(dreq *models.EditDonationReq) (models.DonationDetailResp, error) {
	var d models.Donation
	var dd models.DonationDetail

	isDonationExist, err := r.isDonationExist(dreq.DonationID, dreq.RecipientID)
	if err != nil {
		return models.DonationDetailResp{}, err
	}

	if !isDonationExist {
		return models.DonationDetailResp{}, helper.ErrNoData
	}
	r.DB.Collection("donations").FindOne(context.TODO(), bson.M{"_id": dreq.DonationID}).Decode(&d)
	r.DB.Collection("donation_details").FindOne(context.TODO(), bson.M{"donation_id": dreq.DonationID}).Decode(&dd)
	log.Println("DONATIONS : ", d)
	log.Println("DONATIONS DETAILS: ", dd)

	d.DonationName = dreq.DonationName
	dd.Description = dreq.Description
	dd.DonationType = dreq.DonationType
	dd.Tag = dreq.Tag
	dd.SenderAddress = dreq.SenderAddress
	dd.RelatedLink = dreq.RelatedLink
	dd.Notes = dreq.Notes

	_, err = r.DB.Collection("donations").UpdateOne(
		context.TODO(),
		bson.M{"_id": dreq.DonationID},
		bson.M{"$set": d},
	)

	if err != nil {
		helper.Logging(nil).Error("ERR REPO: ", err)
		return models.DonationDetailResp{}, helper.ErrQuery
	}

	_, err = r.DB.Collection("donation_details").UpdateOne(
		context.TODO(),
		bson.M{"_id": dd.ID},
		bson.M{"$set": d},
	)
	if err != nil {
		helper.Logging(nil).Error("ERR REPO: ", err)
		return models.DonationDetailResp{}, helper.ErrQuery
	}

	resp := models.DonationDetailResp{
		ID:                d.ID,
		RecipientID:       d.RecipientID,
		DonationName:      d.DonationName,
		CreatedAt:         d.CreatedAt,
		Status:            d.Status,
		TargetAmount:      d.TargetAmount,
		AmountCollected:   d.AmountCollected,
		MiscellaneousCost: d.MiscellaneousCost,
		Description:       dd.Description,
		DonationType:      dd.DonationType,
		Tag:               dd.Tag,
		SenderAddress:     dd.SenderAddress,
		RelatedLink:       dd.RelatedLink,
		Notes:             dd.Notes,
	}

	log.Println("donation_detail_resp: ", resp)
	return resp, nil
}

func (r *Repo) DeleteDonation(donation_id string, recipient_id uint) (string, error) {
	d_id, err := primitive.ObjectIDFromHex(donation_id)
	if err != nil {
		helper.Logging(nil).Error("ERR REPO: ", err)
		return "", helper.ErrInvalidId
	}

	isDonationExist, err := r.isDonationExist(d_id, recipient_id)
	if err != nil {
		return "", err
	}

	if !isDonationExist {
		return "", helper.ErrNoData
	}

	res, err := r.DB.Collection("donation_details").DeleteOne(context.TODO(), bson.M{"recipient_id": recipient_id})
	if err != nil {
		helper.Logging(nil).Error("ERR REPO: ", err)
		return "", helper.ErrQuery
	}

	log.Println("DEL DONATION DETAIL ID: ", res)
	res, err = r.DB.Collection("donations").DeleteOne(context.TODO(), bson.M{"_id": d_id})
	if err != nil {
		helper.Logging(nil).Error("ERR REPO: ", err)
		return "", helper.ErrQuery
	}
	log.Println("DEL DONATION ID: ", res)

	out := fmt.Sprintf("Donation with ID: %s deleted successfully", donation_id)
	return out, nil
}
