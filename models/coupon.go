package models

import "gorm.io/gorm"

// Coupon represents a discount coupon
type Coupon struct {
	gorm.Model
	Code       string        `gorm:"unique;not null"`
	Discount   float64       `gorm:"not null"`
	ExpiryDate string        `gorm:"not null"`
	Orders     []OrderCoupon // Many-to-many relationship with Order
}
