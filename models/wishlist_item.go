package models

import "gorm.io/gorm"

// WishlistItem represents an item in a wishlist
type WishlistItem struct {
	gorm.Model
	WishlistID uint `gorm:"not null"` // Foreign key to Wishlist
	ProductID  uint `gorm:"not null"` // Foreign key to Product
	Wishlist   Wishlist
	Product    Product
}
