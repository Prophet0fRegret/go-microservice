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

func (c *Client) UpdateProduct(ctx context.Context, product *models.Product) error {
	productCollection := c.DB.Database("mcr-db").Collection("products")

	filter := bson.M{
		"product_id": product.ProductID,
	}
	updateFields := bson.M{
		"$set": bson.M{
			"name":      product.ProductName,
			"price":     product.Price,
			"vendor_id": product.VendorID,
		},
	}

	result, err := productCollection.UpdateOne(ctx, filter, updateFields)
	if err != nil {
		logrus.WithError(err).Error("Unable to update customer record in Mongo......")
		return err
	}

	logrus.WithFields(logrus.Fields{
		"Matched Count":  result.MatchedCount,
		"Upserted Count": result.UpsertedCount,
	}).Info("Product record updated successfully in Mongo.....")

	return err
}

func (c *Client) DeleteProduct(ctx context.Context, productID string) error {
	productCollection := c.DB.Database("mcr-db").Collection("products")

	filter := bson.M{
		"product_id": productID,
	}

	result, err := productCollection.DeleteOne(ctx, filter)
	if err != nil {
		logrus.WithError(err).Error("Unable to delete product record in Mongo......")
		return err
	}

	logrus.WithFields(logrus.Fields{
		"Deleted Count": result.DeletedCount,
	}).Info("Product record deleted successfully from Mongo.....")

	return err
}
