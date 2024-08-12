package models

import "gorm.io/gorm"

// Invoice represents an invoice for an order
type Invoice struct {
	gorm.Model
	OrderID     uint `gorm:"not null"` // Foreign key to Order
	Order       Order
	InvoiceDate string  `gorm:"not null"`
	TotalAmount float64 `gorm:"not null"`
}
