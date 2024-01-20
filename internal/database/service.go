package database

import (
	"context"

	"github.com/prophet0fregret/go-microservice/internal/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *Client) GetAllServices(ctx context.Context, serviceID string) ([]models.Service, error) {
	var (
		services []models.Service
		filter   = bson.M{}
	)

	if len(serviceID) > 0 {
		filter = bson.M{
			"service_id": serviceID,
		}
	}

	servicesCollection := c.DB.Database("mcr-db").Collection("services")

	cursor, err := servicesCollection.Find(ctx, filter)
	if err != nil {
		logrus.WithError(err).Error("Error occurred while finding customers from collection...")
		return services, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var service models.Service
		err = cursor.Decode(&service)
		if err != nil {
			logrus.WithError(err).Error("Error decoding cursor value...")
			continue
		}

		services = append(services, service)
	}

	return services, nil
}

func (c *Client) CreateNewService(ctx context.Context, service *models.Service) error {
	servicesCollection := c.DB.Database("mcr-db").Collection("services")

	result, err := servicesCollection.InsertOne(ctx, service)
	if err != nil {
		logrus.WithError(err).Error("Unable to insert service record into Mongo......")
		return err
	}

	logrus.WithField("RecordID", result.InsertedID).Info("New service record inserted successfully into Mongo......")

	return err
}
