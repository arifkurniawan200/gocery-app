package app

import (
	"github.com/labstack/echo/v4"
	"grocery-api/internal/model"
	"net/http"
	"strconv"
)

func (u Handler) UpdateProduct(c echo.Context) error {
	var product model.ProductParam
	idParam := c.QueryParam("id")
	id, _ := strconv.Atoi(idParam)
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseFailed{
			Messages: "Failed to parse request body",
			Error:    err.Error(),
		})
	}
	product.ID = int64(id)
	err := u.Product.UpdateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "Failed to get database",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success Update product",
	})
}

func (u Handler) AddProduct(c echo.Context) error {
	var product model.ProductParam

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseFailed{
			Messages: "Failed to parse request body",
			Error:    err.Error(),
		})
	}

	err := u.Product.AddProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "Failed to get database",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success add new product",
	})
}
