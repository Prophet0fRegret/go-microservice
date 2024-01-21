package database

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/prophet0fregret/go-microservice/internal/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ClientInstance DatabaseClient
var initOnce sync.Once

type DatabaseClient interface {
	Ready() bool

	//Customers
	GetAllCustomers(context.Context, string) ([]models.Customer, error)
	CreateNewCustomer(context.Context, *models.Customer) error
	UpdateCustomer(context.Context, *models.Customer) error
	DeleteCustomer(context.Context, string) error

	//Products
	GetAllProducts(context.Context, string) ([]models.Product, error)
	CreateNewProduct(context.Context, *models.Product) error
	UpdateProduct(context.Context, *models.Product) error
	DeleteProduct(context.Context, string) error

	//Services
	GetAllServices(context.Context, string) ([]models.Service, error)
	CreateNewService(context.Context, *models.Service) error
	UpdateService(context.Context, *models.Service) error
	DeleteService(context.Context, string) error

	//Vendors
	GetAllVendors(context.Context, string) ([]models.Vendor, error)
	CreateNewVendor(context.Context, *models.Vendor) error
	UpdateVendor(context.Context, *models.Vendor) error
	DeleteVendor(context.Context, string) error
}

type Client struct {
	DB *mongo.Client
}

func InitDatabaseClient() error {
	var err error
	initOnce.Do(func() {
		ClientInstance, err = NewMongoClient()
	})
	return err
}

func ReturnDatabaseClient() (DatabaseClient, error) {
	var err error
	if ClientInstance == nil {
		err = InitDatabaseClient()
	}
	return ClientInstance, err
}

func NewMongoClient() (DatabaseClient, error) {
	logrus.Info("Initializing new Mongo client.... attempting to connect to Mongo instance")

	connectionString := fmt.Sprintf("mongodb://%s:%s@mongodb:27017/", "admin", "password")
	logrus.Infof("Mongo connection URI - %s", connectionString)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		logrus.WithError(err).Error("Error occurred while connecting to Mongo instance...")
		return nil, err
	}

	logrus.Info("Connected to Mongo instance successfully...")

	newClient := Client{DB: client}

	return &newClient, nil
}

func (c *Client) Ready() bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := c.DB.Ping(ctx, readpref.Primary())
	if err != nil {
		logrus.WithError(err).Error("Mongo instance not pingable...")
		return false
	}

	return true
}
