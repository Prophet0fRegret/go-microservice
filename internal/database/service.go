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

func (c *Client) UpdateService(ctx context.Context, service *models.Service) error {
	serviceCollection := c.DB.Database("mcr-db").Collection("services")

	filter := bson.M{
		"service_id": service.ServiceID,
	}
	updateFields := bson.M{
		"$set": bson.M{
			"name":  service.ServiceName,
			"price": service.Price,
		},
	}

	result, err := serviceCollection.UpdateOne(ctx, filter, updateFields)
	if err != nil {
		logrus.WithError(err).Error("Unable to update service record in Mongo......")
		return err
	}

	logrus.WithFields(logrus.Fields{
		"Matched Count":  result.MatchedCount,
		"Upserted Count": result.UpsertedCount,
	}).Info("Service record updated successfully in Mongo.....")

	return err
}

func (c *Client) DeleteService(ctx context.Context, serviceID string) error {
	serviceCollection := c.DB.Database("mcr-db").Collection("services")

	filter := bson.M{
		"service_id": serviceID,
	}

	result, err := serviceCollection.DeleteOne(ctx, filter)
	if err != nil {
		logrus.WithError(err).Error("Unable to delete service record in Mongo......")
		return err
	}

	logrus.WithFields(logrus.Fields{
		"Deleted Count": result.DeletedCount,
	}).Info("Service record deleted successfully from Mongo.....")

	return err
}
