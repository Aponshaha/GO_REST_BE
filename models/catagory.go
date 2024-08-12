package models

import "gorm.io/gorm"

// Category represents a product category in the database
type Category struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Products    []Product // One-to-many relationship with Product
}
