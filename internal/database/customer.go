package database

import (
	"context"

	"github.com/prophet0fregret/go-microservice/internal/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *Client) GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error) {
	var (
		customers []models.Customer
		filter    = bson.M{}
	)

	if len(emailAddress) > 0 {
		filter = bson.M{
			"email": emailAddress,
		}
	}

	customersCollection := c.DB.Database("mcr-db").Collection("customers")

	cursor, err := customersCollection.Find(ctx, filter)
	if err != nil {
		logrus.WithError(err).Error("Error occurred while finding customers from collection...")
		return customers, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var customer models.Customer
		err = cursor.Decode(&customer)
		if err != nil {
			logrus.WithError(err).Error("Error decoding cursor value...")
			continue
		}

		customers = append(customers, customer)
	}

	return customers, nil
}