package models

import "gorm.io/gorm"

// Product represents a product in the database
type Product struct {
	gorm.Model
	Name          string `gorm:"not null"`
	Description   string
	Price         float64 `gorm:"not null"`
	StockQuantity int     `gorm:"not null"`
	CategoryID    uint    `gorm:"not null"` // Foreign key to Category
	Category      Category
	Images        []ProductImage
	Reviews       []Review
}
