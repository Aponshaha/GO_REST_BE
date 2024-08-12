package models

import "gorm.io/gorm"

// Payment represents a payment made for an order
type Payment struct {
	gorm.Model
	Method string  `gorm:"not null"`
	Amount float64 `gorm:"not null"`
	Order  *Order  `gorm:"foreignKey:PaymentID"` // Use pointer to avoid recursive issues
}
