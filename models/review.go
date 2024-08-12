package models

import "gorm.io/gorm"

// Review represents a review for a product
type Review struct {
	gorm.Model
	ProductID  uint `gorm:"not null"` // Foreign key to Product
	CustomerID uint `gorm:"not null"` // Foreign key to Customer
	Rating     int  `gorm:"not null"` // Rating out of 5
	Comment    string
	Product    Product
	Customer   Customer
}
