package handlers

import (
	"ecommerce-app/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	// "models"
)

// GetProducts fetches all products
func GetProducts(c echo.Context) error {
	var products []models.Product
	if err := DB.Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

// GetProductByID fetches a product by its ID
func GetProductByID(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	if err := DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, product)
}
