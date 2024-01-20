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

func (c *Client) CreateNewCustomer(ctx context.Context, customer *models.Customer) error {
	customersCollection := c.DB.Database("mcr-db").Collection("customers")

	result, err := customersCollection.InsertOne(ctx, customer)
	if err != nil {
		logrus.WithError(err).Error("Unable to insert customer record into Mongo......")
		return err
	}

	logrus.WithField("RecordID", result.InsertedID).Info("New customer record inserted successfully into Mongo......")

	return err
}

func (c *Client) UpdateCustomer(ctx context.Context, customer *models.Customer) error {
	customersCollection := c.DB.Database("mcr-db").Collection("customers")

	filter := bson.M{
		"customer_id": customer.CustomerID,
	}
	updateFields := bson.M{
		"$set": bson.M{
			"first_name": customer.FirstName,
			"last_name":  customer.LastName,
			"email":      customer.EmailID,
			"phone":      customer.Phone,
			"address":    customer.Address,
		},
	}

	result, err := customersCollection.UpdateOne(ctx, filter, updateFields)
	if err != nil {
		logrus.WithError(err).Error("Unable to update customer record in Mongo......")
		return err
	}

	logrus.WithFields(logrus.Fields{
		"Matched Count":  result.MatchedCount,
		"Upserted Count": result.UpsertedCount,
	}).Info("Customer record updated successfully in Mongo.....")

	return err
}
