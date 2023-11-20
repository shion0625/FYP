package domain

import "time"

// represent a model of product.
type Product struct {
	ID            uint      `gorm:"primaryKey;not null"        json:"id"`
	Name          string    `binding:"required,min=3,max=50"   gorm:"not null"   json:"name"`
	Description   string    `binding:"required,min=10,max=100" gorm:"not null"   json:"description"`
	CategoryID    uint      `binding:"omitempty,numeric"       json:"categoryId"`
	Category      Category  `json:"-"`
	BrandID       uint      `gorm:"not null"                   json:"brandId"`
	Brand         Brand     `json:"-"`
	Price         uint      `binding:"required,numeric"        gorm:"not null"   json:"price"`
	DiscountPrice uint      `json:"discountPrice"`
	Image         string    `gorm:"not null"                   json:"image"`
	CreatedAt     time.Time `gorm:"not null"                   json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// this for a specific variant of product.
type ProductItem struct {
	ID            uint      `gorm:"primaryKey;not null" json:"id"`
	ProductID     uint      `binding:"required,numeric" gorm:"not null"  json:"productId"`
	Product       Product   `json:"-"`
	QtyInStock    uint      `binding:"required,numeric" gorm:"not null"  json:"qtyInStock"`
	Price         uint      `binding:"required,numeric" gorm:"not null"  json:"price"`
	SKU           string    `gorm:"unique;not null"     json:"sku"`
	DiscountPrice uint      `json:"discountPrice"`
	CreatedAt     time.Time `gorm:"not null"            json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// for a products category main and sub category as self joining.
type Category struct {
	ID         uint      `gorm:"primaryKey;not null"      json:"-"`
	CategoryID uint      `json:"categoryId"`
	Category   *Category `json:"-"`
	Name       string    `binding:"required,min=1,max=30" gorm:"not null" json:"name"`
}

type Brand struct {
	ID   uint   `gorm:"primaryKey;not null" json:"id"`
	Name string `gorm:"unique;not null"     json:"name"`
}

// variation means size color etc..
type Variation struct {
	ID         uint     `gorm:"primaryKey;not null" json:"-"`
	CategoryID uint     `binding:"required,numeric" gorm:"not null" json:"categoryId"`
	Category   Category `json:"-"`
	Name       string   `binding:"required"         gorm:"not null" json:"name"`
}

// variation option means values are like s,m,xl for size and blue,white,black for Color.
type VariationOption struct {
	ID          uint      `gorm:"primaryKey;not null" json:"-"`
	VariationID uint      `binding:"required,numeric" gorm:"not null" json:"variationId"` // a specific field of variation like color/size
	Variation   Variation `json:"-"`
	Value       string    `binding:"required"         gorm:"not null" json:"value"` // the variations value like blue/XL
}

type ProductConfiguration struct {
	ProductItemID     uint            `gorm:"not null" json:"productItemId"`
	ProductItem       ProductItem     `json:"-"`
	VariationOptionID uint            `gorm:"not null" json:"variationOptionId"`
	VariationOption   VariationOption `json:"-"`
}

// to store a url of productItem Id along a unique url
// so we can ote multiple images url for a ProductItem
// one to many connection.
type ProductImage struct {
	ID            uint        `gorm:"primaryKey;not null" json:"id"`
	ProductItemID uint        `gorm:"not null"            json:"productItemId"`
	ProductItem   ProductItem `json:"-"`
	Image         string      `gorm:"not null"            json:"image"`
}

// offer.
type Offer struct {
	ID           uint      `gorm:"primaryKey;not null"               json:"id"              swaggerignore:"true"`
	Name         string    `binding:"required"                       gorm:"not null;unique" json:"name"`
	Description  string    `binding:"required,min=6,max=50"          gorm:"not null"        json:"description"`
	DiscountRate uint      `binding:"required,numeric,min=1,max=100" gorm:"not null"        json:"discountRate"`
	StartDate    time.Time `binding:"required"                       gorm:"not null"        json:"startDate"`
	EndDate      time.Time `binding:"required,gtfield=StartDate"     gorm:"not null"        json:"endDate"`
}

type OfferProduct struct {
	ID        uint `gorm:"primaryKey;not null" json:"id"`
	OfferID   uint `gorm:"not null"            json:"offerId"`
	Offer     Offer
	ProductID uint `gorm:"not null" json:"productId"`
	Product   Product
}
