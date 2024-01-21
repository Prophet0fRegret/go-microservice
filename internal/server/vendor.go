package server

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/prophet0fregret/go-microservice/internal/models"
	"github.com/sirupsen/logrus"
)

func (e *EchoServer) GetAllVendors(ctx echo.Context) error {
	vendorID := ctx.QueryParam("vendor_id")
	logrus.Info("VendorID Path Param - ", vendorID)

	models, err := e.DB.GetAllVendors(ctx.Request().Context(), vendorID)
	if err != nil {
		logrus.WithError(err).Error("Unable to fetch any vendors from database...")
		ctx.NoContent(http.StatusInternalServerError)
		return err
	}

	return ctx.JSON(http.StatusOK, models)
}

func (e *EchoServer) CreateNewVendor(ctx echo.Context) error {
	var vendor models.Vendor

	if err := ctx.Bind(&vendor); err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unable to unmarshal request body into Vendor model with error : %s", err.Error()))
	}

	vendor.VendorID = uuid.New().String()

	err := e.DB.CreateNewVendor(ctx.Request().Context(), &vendor)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unable to create vendor record in database, error : %s", err.Error()))
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("Record created in database successfully, ID : %s", vendor.VendorID))
}

func (e *EchoServer) UpdateVendor(ctx echo.Context) error {
	var vendor models.Vendor

	if err := ctx.Bind(&vendor); err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unable to unmarshal request body into Vendor model with error : %s", err.Error()))
	}

	err := e.DB.CreateNewVendor(ctx.Request().Context(), &vendor)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unable to update vendor record in database, error : %s", err.Error()))
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("Record update in database successfully, ID : %s", vendor.VendorID))
}

func (e *EchoServer) DeleteVendor(ctx echo.Context) error {
	var vendorID = ctx.Param("vendor_id")

	if err := uuid.Validate(vendorID); err != nil {
		logrus.WithFields(logrus.Fields{
			"VendorID": vendorID,
			"Error":    err,
		}).Error("Invalid vendor ID provided in params....")
		return ctx.String(http.StatusBadRequest, "Invalid vendor ID provided in params....")
	}

	err := e.DB.DeleteVendor(ctx.Request().Context(), vendorID)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unable to delete vendor record in database, error : %s", err.Error()))
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("Record deleted from database successfully, ID : %s", vendorID))
}
