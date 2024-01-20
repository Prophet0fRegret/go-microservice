package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (e *EchoServer) GetAllCustomers(ctx echo.Context) error {
	emailAddress := ctx.QueryParam("emailAddress")
	logrus.Info("Email Path Param - ", emailAddress)

	models, err := e.DB.GetAllCustomers(ctx.Request().Context(), emailAddress)
	if err != nil {
		logrus.WithError(err).Error("Unable to fetch any customers from database...")
		ctx.NoContent(http.StatusInternalServerError)
		return err
	}

	return ctx.JSON(http.StatusOK, models)
}
