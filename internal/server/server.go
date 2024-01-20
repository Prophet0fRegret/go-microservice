package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prophet0fregret/go-microservice/internal/database"
	"github.com/prophet0fregret/go-microservice/internal/models"
	"github.com/sirupsen/logrus"
)

type Server interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}

	server.registerRoutes()
	return server
}

func (s *EchoServer) Start() error {
	err := s.echo.Start(":50051")
	if err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Fatal("Unable to start the echo server......")
	}

	logrus.Info("Echo Server started successfully......")

	return nil
}

func (s *EchoServer) registerRoutes() {
	s.echo.GET("/readiness", s.Readiness)
	s.echo.GET("/liveness", s.Liveness)

	//Customers
	customerGroup := s.echo.Group("/customers")
	customerGroup.GET("/all-customers", s.GetAllCustomers)

	//Products
	productGroup := s.echo.Group("/products")
	productGroup.GET("/all-products", s.GetAllProducts)
}

func (s *EchoServer) Readiness(ctx echo.Context) error {
	ready := s.DB.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, models.HealthCheck{Status: "OK"})
	}
	return ctx.JSON(http.StatusInternalServerError, models.HealthCheck{Status: "FAIL"})
}

func (s *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HealthCheck{Status: "OK"})
}
