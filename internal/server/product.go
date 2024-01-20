package server

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/prophet0fregret/go-microservice/internal/models"
	"github.com/sirupsen/logrus"
)

func (e *EchoServer) GetAllProducts(ctx echo.Context) error {
	vendorID := ctx.QueryParam("vendor_id")
	logrus.Info("VendorID Path Param - ", vendorID)

	models, err := e.DB.GetAllProducts(ctx.Request().Context(), vendorID)
	if err != nil {
		logrus.WithError(err).Error("Unable to fetch any products from database...")
		ctx.NoContent(http.StatusInternalServerError)
		return err
	}

	return ctx.JSON(http.StatusOK, models)
}

func (e *EchoServer) CreateNewProduct(ctx echo.Context) error {
	var product models.Product

	if err := ctx.Bind(&product); err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unable to unmarshal request body into Product model with error : %s", err.Error()))
	}

	product.ProductID = uuid.New().String()

	err := e.DB.CreateNewProduct(ctx.Request().Context(), &product)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unable to create product record in database, error : %s", err.Error()))
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("Record created in database successfully, ID : %s", product.ProductID))
}

func (e *EchoServer) UpdateProduct(ctx echo.Context) error {
	var product models.Product

	if err := ctx.Bind(&product); err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unable to unmarshal request body into Product model with error : %s", err.Error()))
	}

	err := e.DB.UpdateProduct(ctx.Request().Context(), &product)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Unable to update product record in database, error : %s", err.Error()))
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("Record updated in database successfully, ID : %s", product.ProductID))
}
