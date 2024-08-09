package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Cart struct {
	ID              string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	CartItems       []CartItem
	BaseTotalPrice  decimal.Decimal `gorm:"type:decimal(16,2)"`
	TaxAmount       decimal.Decimal `gorm:"type:decimal(16,2)"`
	TaxPercent      decimal.Decimal `gorm:"type:decimal(10,2)"`
	DiscountAmount  decimal.Decimal `gorm:"type:decimal(16,2)"`
	DiscountPercent decimal.Decimal `gorm:"type:decimal(10,2)"`
	GrandTotal      decimal.Decimal `gorm:"type:decimal(16,2)"`
}

func (c *Cart) GetCart(db *gorm.DB, cartID string) (*Cart, error) {
	// var err error
	var cart Cart
	// fmt.Printf("Getting cart with ID: %s\n", cartID)
	// err = db.Debug().Model(Cart{}).Where("id = ?", cartID).First(&cart).Error
	err := db.Debug().Model(&Cart{}).Where("id = ?", cartID).First(&cart).Error
	if err != nil {
		// fmt.Printf("Error getting cart: %v\n", err)
		return nil, err
	}
	// fmt.Printf("Cart found: %+v\n", cart)
	return &cart, nil
}

func (c *Cart) CreateCart(db *gorm.DB, cartID string) (*Cart, error) {
	cart := &Cart{
		ID:              cartID,
		BaseTotalPrice:  decimal.NewFromInt(0),
		TaxAmount:       decimal.NewFromInt(0),
		TaxPercent:      decimal.NewFromFloat(11),
		DiscountAmount:  decimal.NewFromInt(0),
		DiscountPercent: decimal.NewFromInt(0),
		GrandTotal:      decimal.NewFromInt(0),
	}
	// fmt.Printf("Creating cart with ID: %s\n", cartID)
	// err := db.Debug().Create(&cart).Error
	err := db.Debug().Create(cart).Error
	if err != nil {
		// fmt.Printf("Error creating cart: %v\n", err)
		return nil, err
	}
	// fmt.Printf("Cart created: %+v\n", cart)
	return cart, nil
}
