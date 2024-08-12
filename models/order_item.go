package models

import "gorm.io/gorm"

// OrderItem represents an item in an order
type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null"` // Foreign key to Order
	ProductID uint    `gorm:"not null"` // Foreign key to Product
	Quantity  int     `gorm:"not null"`
	UnitPrice float64 `gorm:"not null"`
	Order     Order
	Product   Product
}
