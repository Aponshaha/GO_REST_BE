package models

import "gorm.io/gorm"

// Wishlist represents a customer's wishlist
type Wishlist struct {
	gorm.Model
	CustomerID uint `gorm:"not null"` // Foreign key to Customer
	Customer   Customer
	Items      []WishlistItem
}
