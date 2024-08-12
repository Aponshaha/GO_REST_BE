package models

import "gorm.io/gorm"

// Order represents an order placed by a customer
type Order struct {
	gorm.Model
	OrderNumber string          `gorm:"unique;not null"`
	CustomerID  uint            `gorm:"not null"` // Foreign key to Customer
	Customer    Customer        `gorm:"foreignKey:CustomerID"`
	OrderItems  []OrderItem     `gorm:"foreignKey:OrderID"`
	TotalAmount float64         `gorm:"not null"`
	Status      string          `gorm:"not null"`
	PaymentID   *uint           // Foreign key to Payment (nullable)
	Payment     *Payment        `gorm:"foreignKey:PaymentID"` // Use pointer to avoid recursive issues
	ShippingID  *uint           // Foreign key to ShippingDetail (nullable)
	Shipping    *ShippingDetail `gorm:"foreignKey:ShippingID"` // Use pointer to avoid recursive issues
}
