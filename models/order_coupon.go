package models

import "gorm.io/gorm"

// OrderCoupon represents the usage of a coupon in an order
type OrderCoupon struct {
	gorm.Model
	OrderID  uint `gorm:"not null"` // Foreign key to Order
	CouponID uint `gorm:"not null"` // Foreign key to Coupon
	Order    Order
	Coupon   Coupon
}
