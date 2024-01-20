package server

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/prophet0fregret/go-microservice/internal/models"
	"github.com/sirupsen/logrus"
)

func (e *EchoServer) GetAllServices(ctx echo.Context) error {
	serviceID := ctx.QueryParam("service_id")
	logrus.Info("ServiceID Path Param - ", serviceID)

	models, err := e.DB.GetAllServices(ctx.Request().Context(), serviceID)
	if err != nil {
		logrus.WithError(err).Error("Unable to fetch any services from database...")
		ctx.NoContent(http.StatusInternalServerError)
		return err
	}

	return ctx.JSON(http.StatusOK, models)
}

func (e *EchoServer) CreateNewService(ctx echo.Context) error {
	var service models.Service

	if err := ctx.Bind(&service); err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unable to unmarshal request body into Service model with error : %s", err.Error()))
	}

	service.ServiceID = uuid.New().String()

	err := e.DB.CreateNewService(ctx.Request().Context(), &service)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unable to create service record in database, error : %s", err.Error()))
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("Record created in database successfully, ID : %s", service.ServiceID))
}

func (e *EchoServer) UpdateService(ctx echo.Context) error {
	var service models.Service

	if err := ctx.Bind(&service); err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unable to unmarshal request body into Service model with error : %s", err.Error()))
	}

	err := e.DB.UpdateService(ctx.Request().Context(), &service)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unable to update service record in database, error : %s", err.Error()))
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("Record update in database successfully, ID : %s", service.ServiceID))
}

func (e *EchoServer) DeleteService(ctx echo.Context) error {
	var serviceID = ctx.Param("service_id")

	if err := uuid.Validate(serviceID); err != nil {
		logrus.WithFields(logrus.Fields{
			"ServiceID": serviceID,
			"Error":     err,
		}).Error("Invalid service ID provided in params....")
		return ctx.String(http.StatusBadRequest, "Invalid service ID provided in params....")
	}

	err := e.DB.DeleteService(ctx.Request().Context(), serviceID)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unable to delete service record in database, error : %s", err.Error()))
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("Record deleted from database successfully, ID : %s", serviceID))
}
