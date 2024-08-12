package models

import "gorm.io/gorm"

// Customer represents a customer in the database
type Customer struct {
	gorm.Model
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Email       string `gorm:"unique;not null"`
	PhoneNumber string
	Address     string
	Orders      []Order // One-to-many relationship with Order
}
