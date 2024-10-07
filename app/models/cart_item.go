package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/shopspring/decimal"
)

type CartItem struct {
	ID string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	// CartID          string          `gorm:"size:36;not null"`
	CartID          string          `gorm:"index"`
	ProductID       string          `gorm:"size:36;not null"`
	Qty             int             `gorm:"not null"`
	BasePrice       decimal.Decimal `gorm:"type:decimal(16,2);not null"`
	BaseTotal       decimal.Decimal `gorm:"type:decimal(16,2);not null"`
	TaxAmount       decimal.Decimal `gorm:"type:decimal(16,2);not null"`
	TaxPercent      decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	DiscountAmount  decimal.Decimal `gorm:"type:decimal(16,2);not null"`
	DiscountPercent decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	SubTotal        decimal.Decimal `gorm:"type:decimal(16,2);not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	// ID              string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	// Cart            Cart
	// CartID          string `gorm:"size:36;not null"`
	// Product         Product
	// ProductID       string `gorm:"size:36;not null"`
	// Qty             int
	// BasePrice       decimal.Decimal `gorm:"type:decimal(16,2)"`
	// BaseTotal       decimal.Decimal `gorm:"type:decimal(16,2)"`
	// TaxAmount       decimal.Decimal `gorm:"type:decimal(16,2)"`
	// TaxPercent      decimal.Decimal `gorm:"type:decimal(10,2)"`
	// DiscountAmount  decimal.Decimal `gorm:"type:decimal(16,2)"`
	// DiscountPercent decimal.Decimal `gorm:"type:decimal(10,2)"`
	// SubTotal        decimal.Decimal `gorm:"type:decimal(16,2)"`
	// CreatedAt       time.Time
	// UpdatedAt       time.Time
}

func (c *CartItem) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}

	return nil
}

// func GetCartItem(db *gorm.DB, cartID string, productID string) (*CartItem, error) {
// 	var cartItem CartItem
// 	if err := db.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&cartItem).Error; err != nil {
// 		log.Printf("Error retrieving cart item with CartID %s and ProductID %s: %v", cartID, productID, err)
// 		return nil, err
// 	}
// 	return &cartItem, nil
// }

func GetCartItem(db *gorm.DB, cartID string, productID string) (*CartItem, error) {
	// log.Printf("Fetching CartItem with CartID: %s and ProductID: %s", cartID, productID)
	// var cartItem CartItem
	// if err := db.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&cartItem).Error; err != nil {
	// 	log.Printf("Error retrieving cart item with CartID %s and ProductID %s: %v", cartID, productID, err)
	// 	return nil, err
	// }
	// return &cartItem, nil

	var cartItems []CartItem
	err := db.Where("cart_id = ?", cartID).Find(&cartItems).Error
	if err != nil {
		log.Printf("Error retrieving cart items: %v", err)
	} else {
		log.Printf("Cart items retrieved: %v", cartItems)
	}
	return nil, err
}
