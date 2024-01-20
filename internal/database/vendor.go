package database

import (
	"context"

	"github.com/prophet0fregret/go-microservice/internal/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *Client) GetAllVendors(ctx context.Context, vendorID string) ([]models.Vendor, error) {
	var (
		vendors []models.Vendor
		filter  = bson.M{}
	)

	if len(vendorID) > 0 {
		filter = bson.M{
			"vendor_id": vendorID,
		}
	}

	vendorCollection := c.DB.Database("mcr-db").Collection("vendors")

	cursor, err := vendorCollection.Find(ctx, filter)
	if err != nil {
		logrus.WithError(err).Error("Error occurred while finding customers from collection...")
		return vendors, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var vendor models.Vendor
		err = cursor.Decode(&vendor)
		if err != nil {
			logrus.WithError(err).Error("Error decoding cursor value...")
			continue
		}

		vendors = append(vendors, vendor)
	}

	return vendors, nil
}

func (c *Client) CreateNewVendor(ctx context.Context, vendor *models.Vendor) error {
	vendorsCollection := c.DB.Database("mcr-db").Collection("vendors")

	result, err := vendorsCollection.InsertOne(ctx, vendor)
	if err != nil {
		logrus.WithError(err).Error("Unable to insert vendor record into Mongo......")
		return err
	}

	logrus.WithField("RecordID", result.InsertedID).Info("New vendor record inserted successfully into Mongo......")

	return err
}

func (c *Client) UpdateVendor(ctx context.Context, vendor *models.Vendor) error {
	vendorCollection := c.DB.Database("mcr-db").Collection("vendors")

	filter := bson.M{
		"vendor_id": vendor.VendorID,
	}
	updateFields := bson.M{
		"$set": bson.M{
			"name":    vendor.Name,
			"contact": vendor.Contact,
			"email":   vendor.EmailID,
			"phone":   vendor.Phone,
			"address": vendor.Address,
		},
	}

	result, err := vendorCollection.UpdateOne(ctx, filter, updateFields)
	if err != nil {
		logrus.WithError(err).Error("Unable to update vendor record in Mongo......")
		return err
	}

	logrus.WithFields(logrus.Fields{
		"Matched Count":  result.MatchedCount,
		"Upserted Count": result.UpsertedCount,
	}).Info("Vendor record updated successfully in Mongo.....")

	return err
}
