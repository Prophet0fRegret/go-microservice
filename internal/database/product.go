package database

import (
	"context"

	"github.com/prophet0fregret/go-microservice/internal/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *Client) GetAllProducts(ctx context.Context, vendorID string) ([]models.Product, error) {
	var (
		products []models.Product
		filter   = bson.M{}
	)

	if len(vendorID) > 0 {
		filter = bson.M{
			"vendor_id": vendorID,
		}
	}

	productsCollection := c.DB.Database("mcr-db").Collection("products")

	cursor, err := productsCollection.Find(ctx, filter)
	if err != nil {
		logrus.WithError(err).Error("Error occurred while finding customers from collection...")
		return products, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var product models.Product
		err = cursor.Decode(&product)
		if err != nil {
			logrus.WithError(err).Error("Error decoding cursor value...")
			continue
		}

		products = append(products, product)
	}

	return products, nil
}

func (c *Client) CreateNewProduct(ctx context.Context, product *models.Product) error {
	productsCollection := c.DB.Database("mcr-db").Collection("products")

	result, err := productsCollection.InsertOne(ctx, product)
	if err != nil {
		logrus.WithError(err).Error("Unable to insert product record into Mongo......")
		return err
	}

	logrus.WithField("RecordID", result.InsertedID).Info("New product record inserted successfully into Mongo......")

	return err
}
