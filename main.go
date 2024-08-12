package main

import (
	database "ecommerce-app/Database"
	"log"

	"github.com/labstack/echo/v4"

	// "path/to/your/project/database"
	"path/to/your/project/handlers"

	"github.com/AponShaha/models"
)

func main() {
	database.InitDB()

	// Run migrations
	err := database.DB.AutoMigrate(
		&models.Product{},
		&models.Category{},
		&models.ProductImage{},
		&models.Customer{},
		&models.Order{},
		&models.OrderItem{},
		&models.Payment{},
		&models.Review{},
		&models.Invoice{},
		&models.ShippingDetail{},
		&models.Coupon{},
		&models.OrderCoupon{},
		&models.Wishlist{},
		&models.WishlistItem{},
	)
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	// Initialize the Echo web framework
	e := echo.New()

	// Define routes
	e.GET("/products", handlers.GetProducts)
	e.GET("/products/:id", handlers.GetProductByID)
	// Define other routes...

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
