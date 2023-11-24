package response

import (
	"time"
)

// response for product.
type Product struct {
	ID               uint      `json:"id"`
	CategoryID       uint      `json:"categoryId"`
	Price            uint      `json:"price"`
	DiscountPrice    uint      `json:"discountPrice"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	CategoryName     string    `json:"categoryName"`
	MainCategoryName string    `json:"mainCategoryName"`
	BrandID          uint      `json:"brandId"`
	BrandName        string    `json:"brandName"`
	Image            string    `json:"image"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

// for a specific category representation.
type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// for a specific variation representation.
type Variation struct {
	ID               uint              `json:"id"`
	Name             string            `json:"name"`
	VariationOptions []VariationOption `gorm:"-"`
}

// for a specific variation Value representation.
type VariationOption struct {
	ID    uint   `json:"id"`
	Value string `json:"value"`
}

// for response a specific products all product items.
type ProductItems struct {
	ID               uint                    `json:"id"`
	Name             string                  `json:"name"`
	ProductID        uint                    `json:"productId"`
	Price            uint                    `json:"price"`
	DiscountPrice    uint                    `json:"discountPrice"`
	SKU              string                  `json:"sku"`
	QtyInStock       uint                    `json:"qtyInStock"`
	CategoryName     string                  `json:"categoryName"`
	MainCategoryName string                  `json:"mainCategoryName"`
	BrandID          uint                    `json:"brandId"`
	BrandName        string                  `json:"brandName"`
	VariationValues  []ProductVariationValue `json:"variationValues"`
	Images           []string                `json:"images"`
}

type ProductItemsDB struct {
	ID               uint     `json:"id"`
	Name             string   `json:"name"`
	ProductID        uint     `json:"productId"`
	Price            uint     `json:"price"`
	DiscountPrice    uint     `json:"discountPrice"`
	SKU              string   `json:"sku"`
	QtyInStock       uint     `json:"qtyInStock"`
	CategoryName     string   `json:"categoryName"`
	MainCategoryName string   `json:"mainCategoryName"`
	BrandID          uint     `json:"brandId"`
	BrandName        string   `json:"brandName"`
	Images           []string `json:"images"`
}

type ProductVariationValue struct {
	VariationID       uint   `json:"variationId"`
	Name              string `json:"name"`
	VariationOptionID uint   `json:"variationOptionId"`
	Value             string `json:"value"`
}

// offer response.
type OfferCategory struct {
	OfferCategoryID uint   `json:"offerCategoryId"`
	CategoryID      uint   `json:"categoryId"`
	CategoryName    string `json:"categoryName"`
	DiscountRate    uint   `json:"discountRate"`
	OfferID         uint   `json:"offerId"`
	OfferName       string `json:"offerName"`
}

type OfferProduct struct {
	OfferProductID uint   `json:"offerProductId"`
	ProductID      uint   `json:"productId"`
	ProductName    string `json:"productName"`
	DiscountRate   uint   `json:"discountRate"`
	OfferID        uint   `json:"offerId"`
	OfferName      string `json:"offerName"`
}
