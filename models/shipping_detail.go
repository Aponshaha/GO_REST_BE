package models

import "gorm.io/gorm"

// ShippingDetail represents the shipping details for an order
type ShippingDetail struct {
	gorm.Model
	Address    string `gorm:"not null"`
	City       string `gorm:"not null"`
	PostalCode string `gorm:"not null"`
	Country    string `gorm:"not null"`
	Order      *Order `gorm:"foreignKey:ShippingID"` // Use pointer to avoid recursive issues
}
