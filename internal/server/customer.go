package server

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/prophet0fregret/go-microservice/internal/models"
	"github.com/sirupsen/logrus"
)

func (e *EchoServer) GetAllCustomers(ctx echo.Context) error {
	emailAddress := ctx.QueryParam("emailAddress")
	logrus.Info("Email Path Param - ", emailAddress)

	models, err := e.DB.GetAllCustomers(ctx.Request().Context(), emailAddress)
	if err != nil {
		logrus.WithError(err).Error("Unable to fetch any customers from database...")
		ctx.String(http.StatusInternalServerError, "Failed to fetch customers from the database")
		return err
	}

	return ctx.JSON(http.StatusOK, models)
}

func (e *EchoServer) CreateNewCustomer(ctx echo.Context) error {
	var customer models.Customer

	if err := ctx.Bind(&customer); err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unable to unmarshal request body into Customer model with error : %s", err.Error()))
	}

	customer.CustomerID = uuid.New().String()

	err := e.DB.CreateNewCustomer(ctx.Request().Context(), &customer)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unable to create customer record in database, error : %s", err.Error()))
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("Record created in database successfully, ID : %s", customer.CustomerID))
}
