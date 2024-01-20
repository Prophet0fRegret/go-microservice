package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
