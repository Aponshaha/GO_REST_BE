package models

import "gorm.io/gorm"

// ProductImage represents an image of a product
type ProductImage struct {
	gorm.Model
	URL       string `gorm:"not null"`
	ProductID uint   `gorm:"not null"` // Foreign key to Product
	Product   Product
}
