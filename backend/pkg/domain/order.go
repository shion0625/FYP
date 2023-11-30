package domain

import (
	"time"
)

type ShopOrder struct {
	ID              uint          `gorm:"primaryKey;not null" json:"id"`
	UserID          string        `gorm:"not null"            json:"userId"`
	User            User          `json:"-"`
	OrderDate       time.Time     `gorm:"not null"            json:"orderDate"`
	AddressID       uint          `gorm:"not null"            json:"addressId"`
	Address         Address       `json:"-"`
	OrderTotalPrice uint          `gorm:"not null"            json:"orderTotalPrice"`
	Discount        uint          `gorm:"not null"            json:"discount"`
	PaymentMethodID uint          `json:"paymentMethodId"`
	PaymentMethod   PaymentMethod `json:"-"`
}

type ShopOrderProductItem struct {
	ID            uint `gorm:"primaryKey;unique" json:"id"`
	ShopOrderID   uint `gorm:"not null"          json:"shopOrderId"`
	ShopOrder     ShopOrder
	ProductItemID uint `gorm:"not null" json:"productItemId"`
	ProductItem   ProductItem
	Count         uint `binding:"numeric" json:"count"`
}

type ShopOrderVariation struct {
	ID          uint `gorm:"primaryKey;unique" json:"id"`
	ShopOrderID uint `gorm:"not null"          json:"shopOrderId"`
	ShopOrder
	VariationID       uint `gorm:"not null" json:"variationId"`
	Variation         Variation
	VariationOptionID uint `gorm:"not null" json:"variationOptionId"`
	VariationOption   VariationOption
}
