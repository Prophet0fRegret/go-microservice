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
