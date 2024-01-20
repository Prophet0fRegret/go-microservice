package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (e *EchoServer) GetAllProducts(ctx echo.Context) error {
	vendorID := ctx.QueryParam("vendor_id")
	logrus.Info("VendorID Path Param - ", vendorID)

	models, err := e.DB.GetAllProducts(ctx.Request().Context(), vendorID)
	if err != nil {
		logrus.WithError(err).Error("Unable to fetch any customers from database...")
		ctx.NoContent(http.StatusInternalServerError)
		return err
	}

	return ctx.JSON(http.StatusOK, models)
}
